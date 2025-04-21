package server

import (
	"net/http"

	"github.com/njayp/hades/pkg/power"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		power.Shutdown()
	})
	mux.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		power.Restart()
	})

	// add middleware to all routes
	handler := middlewareLogging(mux)

	return http.ListenAndServe(":7777", handler)
}
