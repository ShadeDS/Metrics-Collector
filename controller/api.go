package controller

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func CreateApi(r *mux.Router, c *Controller) http.Handler {
	r.Handle("/api/v1/metric",
		handlers.LoggingHandler(
			os.Stdout,
			http.HandlerFunc(c.SubmitMetric),
		)).Methods(http.MethodPost)

	return errorHandlerWrapper(r)
}

func errorHandlerWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Printf("Error: %v", err)
				switch err.(type) {
				case string:
					respondWithError(w, http.StatusInternalServerError, err.(string))
				default:
					respondWithError(w, http.StatusInternalServerError, "Unknown error occurred")
				}
				return
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
