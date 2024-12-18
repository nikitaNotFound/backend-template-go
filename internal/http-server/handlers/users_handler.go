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

}

func (h *UsersHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *UsersHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

}
