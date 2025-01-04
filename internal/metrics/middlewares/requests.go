package middlewares

import (
	"app/internal/metrics"
	"net/http"
	"strconv"
	"time"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewCustomResponseWriter(w http.ResponseWriter) *StatusResponseWriter {
	return &StatusResponseWriter{w, http.StatusOK}
}

func (crw *StatusResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func (crw *StatusResponseWriter) Write(b []byte) (int, error) {
	if crw.statusCode == 0 {
		crw.statusCode = http.StatusOK
	}
	return crw.ResponseWriter.Write(b)
}

func RequestsMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		srw := NewCustomResponseWriter(w)
		next.ServeHTTP(srw, r)
		duration := time.Since(start).Seconds()

		metrics.RequestCount.WithLabelValues(r.Method, strconv.Itoa(srw.statusCode), r.URL.Path).Inc()
		metrics.RequestDuration.WithLabelValues(r.Method, strconv.Itoa(srw.statusCode), r.URL.Path).Observe(duration)
	})
}
