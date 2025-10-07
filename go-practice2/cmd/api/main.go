package main

import (
	"log"
	"net/http"

	"go-practice2/handlers"
	"go-practice2/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Setup routes using the new Go 1.22+ pattern matching
	mux.HandleFunc("GET /user", handlers.GetUser)
	mux.HandleFunc("POST /user", handlers.CreateUser)

	// Wrap the mux with middleware
	handler := middleware.AuthMiddleware(mux)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
