package api

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"io"
	"net/http"
)

func createRouter(logWriter io.Writer) (router *mux.Router) {
	router = mux.NewRouter()
	router.Use(
		// Create a logger wrapper
		func(handler http.Handler) http.Handler {
			return handlers.CombinedLoggingHandler(logWriter, handler)
		},
	)

	// When client requests a non-existent path, also log it using the same accessLogWriter, then return 404.
	router.NotFoundHandler = handlers.CombinedLoggingHandler(logWriter, http.NotFoundHandler())

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/version", handleVersion).Methods("GET")

	apiV1 := apiRouter.PathPrefix("/v1").Subrouter()

	// Metrics middleware
	apiV1.Use(func(handler http.Handler) http.Handler {
		return std.Handler("/api/v1", middleware.New(middleware.Config{}))(handler)
	})

	return
}

// A simple HTTP handler to output version data as JSON.
func handleVersion(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	versionData := map[string]string{
		"version": version,
	}

	// Encode the map as JSON and write it to the response
	if err := json.NewEncoder(writer).Encode(versionData); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
