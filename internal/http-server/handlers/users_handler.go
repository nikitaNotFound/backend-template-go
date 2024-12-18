package handlers

import (
	"app/internal/business/usecases/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UsersHandler struct {
	userUseCases *user.UserUseCases
}

func NewUsersHandler(userUseCases *user.UserUseCases) *UsersHandler {
	return &UsersHandler{
		userUseCases: userUseCases,
	}
}

func (h *UsersHandler) RegisterEndpoints(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/login", h.handleLogin)
		r.Post("/register", h.handleRegister)
	})
}

func (h *UsersHandler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *UsersHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
