package api

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	router.NotFoundHandler = handlers.CombinedLoggingHandler(logWriter, http.NotFoundHandler())
	return
}
