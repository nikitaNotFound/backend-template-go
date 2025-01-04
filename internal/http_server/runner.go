package httpserver

import (
	"app/internal/bootstrap"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func StartHttpServer() error {
	deps, err := bootstrap.InitAppDeps()
	if err != nil {
		return err
	}

	httpPort := viper.GetInt("HTTP_PORT")
	if httpPort == 0 {
		httpPort = 8080
	}

	r := chi.NewRouter()
	RegisterRestEndpoints(r, deps)

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
}
