package main

import (
	
	"context"
	
	"fmt"
	"net/http"

	"github.com/machinebox/graphql"
)

// HasuraClient is the client for interacting with Hasura
type HasuraClient struct {
	client *graphql.Client
}

// NewHasuraClient initializes a new Hasura client
func NewHasuraClient(endpoint string) *HasuraClient {
	return &HasuraClient{
		client: graphql.NewClient(endpoint),
	}
}

// InsertUser inserts a user into the Hasura database
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

	// Add headers if necessary (e.g., for admin secret)
	req.Header.Set("x-hasura-admin-secret", "your_admin_secret")

	var resp map[string]interface{}
	err := h.client.Run(ctx, req, &resp)
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}

	fmt.Println("Inserted user:", resp)
	return nil
}
