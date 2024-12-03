package main

import (
	"context"
	// "crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
	 "golang.org/x/crypto/bcrypt"
)
// Struct for Login Request
type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Struct for Login Response
type LoginResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userId,omitempty"` // Include only on success
}
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
	Image       string `json:"image"`  // Assume base64 or image URL
	UserID      string `json:"userId"` // Ensure authenticated user sends this
}

// Hasura GraphQL client configuration
var hasuraEndpoint = "http://localhost:8080/v1/graphql" // Adjust to match your Hasura instance
var hasuraAdminSecret = "myadminsecretkey"              // Replace with your Hasura admin secret (if set)












// Middleware to check if the user exists in the database
func CheckUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Assume the Authorization header contains a user ID (e.g., "Bearer <user-id>")
		parts := strings.Split(authHeader, " ")
if len(parts) != 2 || parts[0] != "Bearer" {
    http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
    return
}
userID := parts[1] // Expect userID or decode JWT based on your setup


		// Query Hasura to check if the user exists
		client := graphql.NewClient(hasuraEndpoint)
		req := graphql.NewRequest(`
			query($id: String!) {
				users_by_pk(id: $id) {
					id
				}
			}
		`)
		req.Var("id", userID)
		req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

		var resp struct {
			User struct {
				ID string `json:"id"`
			} `json:"users_by_pk"`
		}
		if err := client.Run(context.Background(), req, &resp); err != nil || resp.User.ID == "" {
			http.Error(w, "User not authorized", http.StatusUnauthorized)
			return
		}

		// User exists, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// Example protected handler
func AuthorizedPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the authorized page!",
	})
}










func handleCreateRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var recipe RecipeInput
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate input (e.g., required fields)
	if recipe.Title == "" || recipe.Description == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Simulate recipe creation and return a response
	response := map[string]string{
		"message": "Recipe created successfully!",
		"recipeId": "example-id-123", // Replace with actual ID from your database
	}
	json.NewEncoder(w).Encode(response)
}

// RecipeHandler handles recipe creation
func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	var input RecipeInput

	// Set response headers for CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Match your frontend origin
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle OPTIONS preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate input
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
		"message":  "Recipe created successfully!",
		"recipeId": resp.InsertRecipesOne.ID,
	})
}

// SignupHandler for user signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input SignupInput

	// Decode input JSON
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create a GraphQL client
	client := graphql.NewClient(hasuraEndpoint)

	// Set up the GraphQL mutation
	query := `
		mutation($username: String!, $email: String!, $password: String!) {
			insert_users_one(object: {username: $username, email: $email, password: $password}) {
				id
			}
		}
	`

	// Create a new request and set variables
	req := graphql.NewRequest(query)
	req.Var("username", input.Username)
	req.Var("email", input.Email)
	req.Var("password", string(hashedPassword))

	// Add Hasura admin secret header (if required)
	req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

	// Execute the GraphQL request
	var resp struct {
		InsertUsersOne struct {
			ID string `json:"id"`
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



func LoginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var input LoginInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input format", http.StatusBadRequest)
        return
    }

    // Log the input email and password for debugging
    log.Printf("Login attempt for email: %s", input.Email)

    // Fetch user by email
    emailQuery := `
    query($email: String!) {
        users(where: {email: {_eq: $email}}) {
            id
            password
        }
    }`
    client := graphql.NewClient(hasuraEndpoint)
    req := graphql.NewRequest(emailQuery)
    req.Var("email", input.Email)
    req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

    var resp struct {
        Users []struct {
            ID       string `json:"id"`
            Password string `json:"password"`
        } `json:"users"`
    }

    if err := client.Run(context.Background(), req, &resp); err != nil {
        http.Error(w, "Failed to process login request", http.StatusInternalServerError)
        return
    }

    if len(resp.Users) == 0 {
        log.Println("Email not found:", input.Email)
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(LoginResponse{
            Message: "Email not registered",
        })
        return
    }

    // Log the retrieved user data for debugging
    log.Printf("User found: ID: %s", resp.Users[0].ID)

    // Compare password
    log.Println("Comparing passwords")
    err := bcrypt.CompareHashAndPassword([]byte(resp.Users[0].Password), []byte(input.Password))
    if err != nil {
        log.Println("Invalid password:", err)
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(LoginResponse{
            Message: "Invalid password",
        })
        return
    }

    // Success
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(LoginResponse{
        Message: "Login successful",
        UserID:  resp.Users[0].ID,
    })
}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/create-recipe", RecipeHandler).Methods("POST", "OPTIONS")
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	
// 	// CORS configuration
// 	corsHandler := handlers.CORS(
// 		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Adjust based on your front-end origin
// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
// 		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
// 	)

// 	log.Println("Server is running on http://localhost:8081")
// 	log.Fatal(http.ListenAndServe(":8081", corsHandler(r)))
// }
// Configure CORS


// Protected routes
protected := r.PathPrefix("/authorized").Subrouter()
protected.Use(CheckUserMiddleware)
protected.HandleFunc("/page", AuthorizedPageHandler).Methods("GET")




corsOptions := handlers.CORS(
	handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Allow your frontend origin
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), // Allow necessary methods
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allow necessary headers
)

// Start the server with CORS middleware
log.Println("Server running on port 8081")
log.Fatal(http.ListenAndServe(":8081", corsOptions(r)))

}

