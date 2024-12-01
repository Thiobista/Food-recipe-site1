package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
)

// Struct for input from the request
type SignupInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct for response to the client
type SignupResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userId"` // Changed from int to string
}
// RecipeInput struct for receiving recipe data
type RecipeInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Steps       string `json:"steps"`
	Time        string `json:"time"`
	Category    string `json:"category"`
	Image       string `json:"image"` // Assume base64 or image URL
	UserID      string `json:"userId"` // Ensure authenticated user sends this
}

// RecipeHandler handles recipe creation
func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input RecipeInput

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Example: Validate input
	if input.Title == "" || input.Description == "" {
		http.Error(w, "Title and Description are required", http.StatusBadRequest)
		return
	}

	// Create a GraphQL client
	client := graphql.NewClient(hasuraEndpoint)

	// Define the GraphQL mutation
	req := graphql.NewRequest(`
		mutation($title: String!, $description: String!, $ingredients: String!, $steps: String!, $time: String!, $category: String!, $userId: String!) {
			insert_recipes_one(object: {title: $title, description: $description, ingredients: $ingredients, steps: $steps, time: $time, category: $category, user_id: $userId}) {
				id
			}
		}
	`)

	// Set mutation variables
	req.Var("title", input.Title)
	req.Var("description", input.Description)
	req.Var("ingredients", input.Ingredients)
	req.Var("steps", input.Steps)
	req.Var("time", input.Time)
	req.Var("category", input.Category)
	req.Var("userId", input.UserID)

	// Add Hasura admin secret
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	// Execute the mutation
	var resp struct {
		InsertRecipesOne struct {
			ID string `json:"id"`
		} `json:"insert_recipes_one"`
	}

	err = client.Run(context.Background(), req, &resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting recipe: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Recipe created successfully!",
		"recipeId": resp.InsertRecipesOne.ID,
	})
}

// Hasura GraphQL client configuration
var hasuraEndpoint = "http://localhost:8080/v1/graphql" // Adjust to match your Hasura instance
var hasuraAdminSecret = "myadminsecretkey"              // Replace with your Hasura admin secret (if set)

// SignupHandler for user signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input SignupInput

	// Decode input JSON
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash the password
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(input.Password)))

	// Create a GraphQL client
	client := graphql.NewClient(hasuraEndpoint)

	// Set up the GraphQL mutation
	req := graphql.NewRequest(`
		mutation($username: String!, $email: String!, $password: String!) {
			insert_users_one(object: {username: $username, email: $email, password: $password}) {
				id
			}
		}
	`)

	// Set variables for the mutation
	req.Var("username", input.Username)
	req.Var("email", input.Email)
	req.Var("password", passwordHash)

	// Add Hasura admin secret header (if required)
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	// Execute the GraphQL request
	var resp struct {
		InsertUsersOne struct {
			ID string `json:"id"` // Adjusted type to string
		} `json:"insert_users_one"`
	}

	err = client.Run(context.Background(), req, &resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	response := SignupResponse{
		Message: "Signup successful!",
		UserID:  resp.InsertUsersOne.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/create-recipe", RecipeHandler).Methods("POST")

	// CORS configuration
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Adjust based on your front-end origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", corsHandler(r)))
}
