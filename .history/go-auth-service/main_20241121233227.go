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
	return &HasuraClient{
		client: graphql.NewClient(endpoint),
	}
}

// InsertUser inserts a new user into the Hasura database.
func (h *HasuraClient) InsertUser(ctx context.Context, username, password string) error {
	req := graphql.NewRequest(`
		mutation($username: String!, $password: String!) {
			insert_users_one(object: {username: $username, password: $password}) {
				id
			}
		}
	`)
	req.Var("username", username)
	req.Var("password", password)

	var resp map[string]interface{}
	err := h.client.Run(ctx, req, &resp)
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
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

	usersDB[user.Username] = user.Password

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
type UploadResponse struct {
    URL string `json:"url"`
}

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Failed to read file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Save file locally
    outFile, err := os.Create("./uploads/" + header.Filename)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
    defer outFile.Close()
    io.Copy(outFile, file)

    response := UploadResponse{
        URL: fmt.Sprintf("http://localhost:8080/uploads/%s", header.Filename),
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/signup", Signup).Methods("POST")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/protected-route", ProtectedRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}
