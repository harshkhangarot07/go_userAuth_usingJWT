package main

import (
	"go_userAuth/internal/handlers"
	"go_userAuth/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if the environment variable is not set
	}
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Apply middleware to the protected route
	r.Handle("/protected", middleware.JWTMiddleware(http.HandlerFunc(handlers.ProtectedHandler))).Methods("GET")
	log.Printf("Server is starting on port %s...", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
