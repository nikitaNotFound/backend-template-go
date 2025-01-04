package handlers

import (
	"app/internal/bootstrap"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	*bootstrap.AppDeps
}

func NewAuthHandler(deps *bootstrap.AppDeps) *AuthHandler {
	return &AuthHandler{
		AppDeps: deps,
	}
}

func (h *AuthHandler) RegisterEndpoints(r chi.Router) {
	r.Post("/login", h.handleLogin)
	r.Post("/register", h.handleRegister)
}

func (h *AuthHandler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
