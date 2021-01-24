package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("bohemian.reppysody")

// User struct
type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response struct
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JwtToken struct
type JwtToken struct {
	Token string `json:"token"`
}

func (app *App) profile(w http.ResponseWriter, r *http.Request) {
	token := context.Get(r, "decoded")
	fmt.Println(token)
}

func (app *App) signup(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user User
	var response Response

	json.Unmarshal(body, &user)
	hash, _ := HashPassword(user.Password)

	result, err := app.Database.Exec("INSERT INTO `users` (email, password) VALUES (?, ?)", user.Email, hash)
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	userID, err := result.LastInsertId()

	response.Status = http.StatusCreated
	response.Message = http.StatusText(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("ID", strconv.FormatInt(userID, 10))

	json.NewEncoder(w).Encode(response)
}

func (app *App) signin(w http.ResponseWriter, r *http.Request) {
	var user User
	var response Response

	json.NewDecoder(r.Body).Decode(&user)
	password := user.Password

	query := `SELECT id, email, password FROM users WHERE email = ?`
	err := app.Database.QueryRow(query, user.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		response.Status = http.StatusNotFound
		response.Message = http.StatusText(http.StatusNotFound)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)

		return
	}

	match := CheckPasswordHash(password, user.Password)
	if !match {
		response.Status = http.StatusNotFound
		response.Message = http.StatusText(http.StatusNotFound)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Email,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = http.StatusText(http.StatusInternalServerError)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)

		return
	}

	response.Status = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
	response.Data = &JwtToken{
		Token: tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	return
}

// HashPassword function
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash function
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
