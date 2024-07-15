package auth

import (
    "context"
    "log"

    "data-definitions/auth" // Import generated protobuf code
    "data-definitions/user" // Import generated protobuf code
    "google.golang.org/grpc"
)


// Auth data service client 
type AuthClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(client pb.AuthServiceClient) *AuthClient {
    return &AuthClient{
       client: client,
    }
}

// Save the user details -Call CreateUser from auth_data_service
func (c *AuthClient) SaveUser(ctx context.Context, user *user.User) error {
    _, err := c.client.CreateUser(ctx, &auth.SaveUserRequest{User: user})
    if err != nil {
        log.Printf("[ERROR] Failed to save user via gRPC: %v", err)
        return err
    }
    return nil
}

// gets the user details -Call GetUserByUsername from auth_data_service
func (c *AuthServiceClient) GetUserByUsername(ctx context.Context, username string) (*User, error) {
    user, err := c.client.GetUserByUsername(ctx, &GetUserByUsernameRequest{Username: username})
    if err != nil {
        log.Printf("[ERROR] Failed to get user by username: %v", err)
        return nil, err
    }

    return user, nil
}