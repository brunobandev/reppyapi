package handler

import (
	"net/http"

	"github.com/brunobandev/reppyapi/internal/service"
	"github.com/go-chi/chi"
)

type PropertyHandler struct {
	propertyService service.PropertyService
}

func NewPropertyHandler(ps service.PropertyService) *PropertyHandler {
	return &PropertyHandler{ps}
}

func (b *PropertyHandler) Route(r chi.Router) {
	r.Get("/hello", b.hello)
}

func (b *PropertyHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
