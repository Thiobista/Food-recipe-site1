                   package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Structs for inputs and responses
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userId,omitempty"`
	Token   string `json:"token,omitempty"`
}

type SignupInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

var hasuraEndpoint = "http://localhost:8080/v1/graphql"
var hasuraAdminSecret = "myadminsecretkey"
var jwtSecretKey = "your_secret_key"

// Middleware for token validation
func CheckUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}





func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hashed password with its plain text version.
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }






// Hashes the password using bcrypt
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return "", err
    }
    return string(hashedPassword), nil
}

// Signup handler
func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var input SignupInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input format", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := hashPassword(input.Password)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        http.Error(w, "Error processing password", http.StatusInternalServerError)
        return
    }

    // Prepare GraphQL mutation
    client := graphql.NewClient(hasuraEndpoint)
    req := graphql.NewRequest(`
    mutation CreateUser($username: String!, $email: String!, $hashedPassword: String!) {
        insert_users_one(object: {username: $username, email: $email, password: $hashedPassword}) {
            id
            username
            email
        }
    }
    `)
    req.Var("username", input.Username)
    req.Var("email", input.Email)
    req.Var("hashedPassword", hashedPassword) // Use the hashed password
    req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

    var resp struct {
        InsertUsersOne struct {
            ID string `json:"id"`
        } `json:"insert_users_one"`
    }

    // Execute the mutation
    if err := client.Run(context.Background(), req, &resp); err != nil {
        log.Printf("Error executing GraphQL mutation: %v", err)
        http.Error(w, "Error creating user account", http.StatusInternalServerError)
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

// CheckPasswordHash compares a bcrypt hashed password with its plain text version.
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// Login handler
func handleLogin(w http.ResponseWriter, r *http.Request) {
    var loginInput LoginInput
    if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
        http.Error(w, "Invalid input format", http.StatusBadRequest)
        return
    }

    // Query user data
    client := graphql.NewClient(hasuraEndpoint)
    req := graphql.NewRequest(`
        query($email: String!) {
            users(where: {email: {_eq: $email}}) {
                id
                password
            }
        }
    `)
    req.Var("email", loginInput.Email)
    req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

    var resp struct {
        Users []struct {
            ID       string `json:"id"`
            Password string `json:"password"`
        } `json:"users"`
    }

    if err := client.Run(context.Background(), req, &resp); err != nil {
        log.Printf("Error executing GraphQL query: %v", err)
        http.Error(w, "Error finding user", http.StatusInternalServerError)
        return
    }

    if len(resp.Users) == 0 {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    // Check password hash
    user := resp.Users[0]
    if !CheckPasswordHash(loginInput.Password, user.Password) {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    // Respond with success
    response := LoginResponse{
        Message: "Login successful!",
        UserID:  user.ID,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}



func main() {


	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/login", handleLogin).Methods("POST")
	r.HandleFunc("/signup", SignupHandler).Methods("POST")

	// CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", corsOptions(r)))
}
