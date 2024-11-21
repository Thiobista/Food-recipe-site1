// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	 "golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var usersDB = map[string]string{
	// Example users (username:password)
	"testuser": "password123",
}

// Signup endpoint - for creating a new user (simplified version)
func Signup(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save user logic here (e.g., store in database)
	usersDB[user.Username] = user.Password

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
func ValidateTokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
            return
        }

        token, err := ValidateJWT(tokenStr)
        if err != nil || !token.Valid {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// Login endpoint - for generating JWT after successful authentication
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if user exists and password matches
	storedPassword, exists := usersDB[user.Username]
	if !exists || storedPassword != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT for user
	token, err := GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return token to the user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8081", r))
}
