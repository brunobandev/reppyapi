package app

import (
	"fmt"
	"net/http"

	"github.com/brunobandev/reppyapi/model"
)

func (app *App) getProperty(w http.ResponseWriter, r *http.Request) {
	userID := 1
	var property model.Property
	property = model.GetPropertyByUserID(app.Database, userID)

	fmt.Println(property, userID)
}

func (app *App) updateProperty(w http.ResponseWriter, r *http.Request) {

}
