package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// App struct
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

// SetupRouter :)
func (app *App) SetupRouter() {
	app.Router.Methods("POST").Path("/join").HandlerFunc(app.signup)
	app.Router.Methods("POST").Path("/login").HandlerFunc(app.signin)
	app.Router.Methods("GET").Path("/profile").HandlerFunc(app.middleware(app.profile))

	app.Router.Methods("POST").Path("/property").HandlerFunc(app.getProperty)
}

func (app *App) middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return jwtKey, nil
				})

				if err != nil {
					json.NewEncoder(rw).Encode(Response{Message: err.Error()})
					return
				}

				if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					context.Set(r, "decoded", token.Claims)
					next(rw, r)
				} else {
					json.NewEncoder(rw).Encode(Response{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(rw).Encode(Response{Message: "An authorization header is required"})
		}
	})
}
