package httpserver

import (
	"app/internal/bootstrap"
	"app/internal/http-server/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRestEndpoints(r *chi.Mux, deps *bootstrap.AppDeps) {
	authHandler := handlers.NewAuthHandler(deps)
	dataHandler := handlers.NewDataHandler(deps)

	r.Route("/users", authHandler.RegisterEndpoints)
	r.Route("/data", dataHandler.RegisterEndpoints)
}
