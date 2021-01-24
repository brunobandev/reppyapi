package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	serverName := "localhost:33060"
	user := "reppy"
	password := "userpassword"
	dbName := "reppy"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
