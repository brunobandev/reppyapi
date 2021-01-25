package handler

import (
	"log"
	"net/http"

	"github.com/brunobandev/reppyapi/internal/container"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type handlers struct {
	Property *PropertyHandler
}

func initHandlers(c container.Container) handlers {
	return handlers{
		Property: NewPropertyHandler(c.Services.PropertyService),
	}
}

func StartServer(c container.Container) {
	log.Println("Initializing the Http Server at localhost:8080")
	handlers := initHandlers(c)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/properties", handlers.Property.Route)
	})

	http.ListenAndServe(":8080", r)
}
