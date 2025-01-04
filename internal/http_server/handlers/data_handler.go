package handlers

import (
	"app/internal/bootstrap"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DataHandler struct {
	*bootstrap.AppDeps
}

func NewDataHandler(deps *bootstrap.AppDeps) *DataHandler {
	return &DataHandler{
		AppDeps: deps,
	}
}

func (h *DataHandler) RegisterEndpoints(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/data", h.handleGetData)
	})
}

func (h *DataHandler) handleGetData(w http.ResponseWriter, r *http.Request) {}
