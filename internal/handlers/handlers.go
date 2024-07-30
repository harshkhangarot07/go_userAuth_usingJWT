package handlers

import (
	"encoding/json"
	"fmt"
	"go_userAuth/internal/auth"
	"go_userAuth/internal/middleware"
	"go_userAuth/internal/models"
	"io/ioutil"

	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Request Body:", string(body))

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println("Parsed User:", user)

	err = auth.RegisterUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if !auth.AuthenticateUser(user.Username, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token":   token,
		"message": "Login successful",
	})
}
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.ClaimsContextKey).(*auth.Claims)
	if !ok {
		http.Error(w, "No claims in context", http.StatusUnauthorized)
		return
	}

	fmt.Println("Claims in Handler:", claims) // Debugging statement

	username := claims.Username
	fmt.Fprintf(w, "This is a protected route! Welcome %s", username)
}
