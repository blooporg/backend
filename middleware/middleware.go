package middleware

import (
	"log"
	"net/http"
)

type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorHandling handles the error and returns http.Handler
func ErrorHandling(h MyHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			// Instead of terminating the program, let's just log the error
			log.Printf("An error occurred: %v\n", err)

			// And return an HTTP 500 status code
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
