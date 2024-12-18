package handlers

import (
	"app/internal/business/usecases/data"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DataHandler struct {
	dataUseCases *data.DataUseCases
}

func NewDataHandler(dataUseCases *data.DataUseCases) *DataHandler {
	return &DataHandler{
		dataUseCases: dataUseCases,
	}
}

func (h *DataHandler) RegisterEndpoints(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/data", h.handleGetData)
	})
}

func (h *DataHandler) handleGetData(w http.ResponseWriter, r *http.Request) {}
