package httpserver

import (
	"app/internal/bootstrap"
	"app/internal/metrics"
	"app/internal/metrics/middlewares"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	metrics.InitMetrics()

	r := chi.NewRouter()
	r.Use(middlewares.RequestsMetricsMiddleware)
	r.Handle("/metrics", promhttp.Handler())
	RegisterRestEndpoints(r, deps)

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
}
