package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log method and path
		log.Printf("%s %s", r.Method, r.URL.Path)

		// Check for X-API-Key header
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "unauthorized"})
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
