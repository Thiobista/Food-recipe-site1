package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var usersDB = map[string]string{
	"testuser": "password123",
}

type HasuraClient struct {
	client *graphql.Client
}

// NewHasuraClient creates a new Hasura GraphQL client.
func NewHasuraClient(endpoint string) *HasuraClient {
	return &HasuraClient{endpoint: endpoint}
}

func (h *HasuraClient) InsertUser(ctx context.Context, username, password string) error {
	query := `
		mutation($username: String!, $password: String!) {
			insert_users_one(object: {username: $username, password: $password}) {
				id
			}
		}
	`

	variables := map[string]interface{}{
		"username": username,
		"password": password,
	}

	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to encode payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", h.endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", "your_admin_secret") // Replace with your Hasura admin secret

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// Signup endpoint
func Signup(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hasuraClient := NewHasuraClient("http://localhost:8080/v1/graphql")

	err = hasuraClient.InsertUser(context.Background(), user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error inserting user to Hasura: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// Login endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check user credentials
	storedPassword, exists := usersDB[user.Username]
	if !exists || storedPassword != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// ProtectedRoute endpoint
func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the protected route!"))
}

// Default home handler
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/protected-route", ProtectedRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}
