package api

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
		func(handler http.Handler) http.Handler {
			return handlers.CombinedLoggingHandler(logWriter, handler)
		},
	)

	router.NotFoundHandler = handlers.CombinedLoggingHandler(logWriter, http.NotFoundHandler())

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiV1 := apiRouter.PathPrefix("/v1").Subrouter()

	return
}
