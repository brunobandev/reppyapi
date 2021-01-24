package main

import (
	"log"
	"net/http"

	"github.com/brunobandev/reppyapi/app"
	"github.com/brunobandev/reppyapi/db"
	"github.com/gorilla/mux"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
