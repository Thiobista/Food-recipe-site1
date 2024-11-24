package main

import (
	// "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"github.com/gorilla/mux"
	// "github.com/machinebox/graphql"
	"crypto/sha256"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var usersDB = map[string]string{
	"testuser": "password123",
}

// type HasuraClient struct {
// 	client *graphql.Client
// }

type SignupInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	UserID  int    `json:"userId"`
}

// NewHasuraClient creates a new Hasura GraphQL client.
// func NewHasuraClient(endpoint string) *HasuraClient {
// 	return &HasuraClient{
// 		client: graphql.NewClient(endpoint),
// 	}
// }

// InsertUser inserts a new user into the Hasura database.
// func (h *HasuraClient) InsertUser(ctx context.Context, username, password string) error {
// 	req := graphql.NewRequest(`
// 		mutation($username: String!, $password: String!) {
// 			insert_users_one(object: {username: $username, password: $password}) {
// 				id
// 			}
// 		}
// 	`)
// 	req.Var("username", username)
// 	req.Var("password", password)

// 	var resp map[string]interface{}
// 	err := h.client.Run(ctx, req, &resp)
// 	if err != nil {
// 		return fmt.Errorf("error inserting user: %v", err)
// 	}

// 	return nil
// }

// SignupHandler is the handler for the signup endpoint.
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Input SignupInput `json:"input"`
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash the password
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(input.Input.Password)))

	// Insert user into the database (mocked example)
	// Replace this mock with actual database interaction
	fmt.Printf("Storing user: Username=%s, Email=%s, PasswordHash=%s\n", 
		input.Input.Username, input.Input.Email, passwordHash)

	// Example User ID after successful insertion
	userID := 1 

	// Return success response
	response := SignupResponse{
		Message: "Signup successful!",
		UserID:  userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Login is the login handler endpoint.
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

// ProtectedRoute is the endpoint for accessing a protected route.
func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the protected route!"))
}

// Home is the default home handler.
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}

type UploadResponse struct {
    URL string `json:"url"`
}

// uploadImageHandler handles image upload.
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
	r.HandleFunc("/signup", SignupHandler).Methods("POST") // Updated this handler
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/protected-route", ProtectedRoute).Methods("GET")
	http.HandleFunc("/upload", uploadImageHandler)
	r.HandleFunc("/signup", SignupHandler).Methods("POST")


	log.Fatal(http.ListenAndServe(":8081", r))
}
