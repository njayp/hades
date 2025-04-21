package server

import (
	"log/slog"
	"net/http"
)

func middlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log request
		slog.Info("request received", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
