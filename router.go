package main

import (
	"github.com/carolineeey/multifinance-api/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func createRouter(apiClient *api.Client, logWriter io.Writer) (router *mux.Router) {
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
	apiV1 := apiRouter.PathPrefix("/v1").Subrouter()

	customerV1 := apiV1.PathPrefix("/customer").Subrouter()
	customerV1.HandleFunc("/get", apiClient.HandleGetCustomer).Methods("GET")
	customerV1.HandleFunc("/add", apiClient.HandleAddCustomer).Methods("POST")
	customerV1.HandleFunc("/edit", apiClient.HandleEditCustomer).Methods("POST")
	customerV1.HandleFunc("/delete", apiClient.HandleDeleteCustomer).Methods("DELETE")

	return
}
