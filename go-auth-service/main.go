package main

import (
	"context"
	"encoding/json"
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

type RecipeInput struct {
    RecipeID        string   `json:"recipe_id"` // Add this line if it's needed for associating recipe with ingredients/steps
    Title           string   `json:"title"`
    Description     string   `json:"description"`
    CategoryID      string   `json:"category_id"`
    TimeToPrepare   int      `json:"time_to_prepare"`
    FeaturedImageURL string  `json:"featured_image_url"`
    UserID          string   `json:"user_id"`
    Ingredients     []IngredientInput `json:"ingredients"`
    Steps           []StepInput      `json:"steps"`
}


type IngredientInput struct {
    IngredientName string `json:"ingredient_name"`
    Quantity       string `json:"quantity"`
}

type StepInput struct {
    StepNumber     int    `json:"step_number"`
    StepDescription string `json:"step_description"`
}


type RecipeResponse struct {
	Message string `json:"message"`
	RecipeID string `json:"recipeId"`
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

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": user.ID,
    })
    tokenString, err := token.SignedString([]byte(jwtSecretKey))
    if err != nil {
        log.Printf("Error signing token: %v", err)
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Respond with success and token
    response := LoginResponse{
        Message: "Login successful!",
        UserID:  user.ID,
        Token:   tokenString,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// CreateRecipeHandler handles the creation of new recipes
// CreateRecipeHandler handles the creation of new recipes
func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
    var input RecipeInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input format", http.StatusBadRequest)
        return
    }

    // Validate inputs
    if input.Title == "" || input.Description == "" || input.UserID == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

    // Dynamically map ingredients and steps
    ingredients := make([]map[string]interface{}, len(input.Ingredients))
    for i, ing := range input.Ingredients {
        ingredients[i] = map[string]interface{}{
            "recipe_id":       input.RecipeID,
            "ingredient_name": ing.IngredientName,
            "quantity":        ing.Quantity,
        }
    }

    steps := make([]map[string]interface{}, len(input.Steps))
    for i, step := range input.Steps {
        steps[i] = map[string]interface{}{
            "recipe_id":         input.RecipeID,
            "step_number":       step.StepNumber,
            "step_description":  step.StepDescription,
        }
    }

    // Prepare GraphQL mutation
    client := graphql.NewClient(hasuraEndpoint)
    req := graphql.NewRequest(`
    mutation CreateRecipe($input: recipes_insert_input!, $ingredients: [ingredients_insert_input!]!, $steps: [steps_insert_input!]!) {
        insert_recipes_one(object: $input) {
            id
            title
            description
        }
        insert_ingredients(objects: $ingredients) {
            returning {
                id
                ingredient_name
                quantity
            }
        }
        insert_steps(objects: $steps) {
            returning {
                id
                step_number
                step_description
            }
        }
    }`)
    req.Var("input", map[string]interface{}{
        "title":             input.Title,
        "description":       input.Description,
        "category_id":       input.CategoryID,
        "user_id":           input.UserID,
        "time_to_prepare":   input.TimeToPrepare,
        "featured_image_url": input.FeaturedImageURL,
    })
    req.Var("ingredients", ingredients)
    req.Var("steps", steps)
    req.Header.Set("x-hasura-admin-secret", hasuraAdminSecret)

    var resp struct {
        InsertRecipesOne struct {
            ID string `json:"id"`
        } `json:"insert_recipes_one"`
    }

    if err := client.Run(context.Background(), req, &resp); err != nil {
        log.Printf("Error inserting recipe: %v", err)
        http.Error(w, "Error creating recipe", http.StatusInternalServerError)
        return
    }

    // Return response
    response := RecipeResponse{
        Message: "Recipe created successfully!",
        RecipeID: resp.InsertRecipesOne.ID,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	// Public routes 
	r.HandleFunc("/login", handleLogin).Methods("POST")
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
    r.Handle("/api/recipes", CheckUserMiddleware(http.HandlerFunc(CreateRecipeHandler))).Methods("POST")

	// CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", corsOptions(r)))
}
