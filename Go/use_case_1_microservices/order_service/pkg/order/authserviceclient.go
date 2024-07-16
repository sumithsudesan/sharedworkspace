package order

import (
    "context"

    "google.golang.org/grpc"
    pbAuth "data-definitions/auth"
)

// Auth client
type authServiceGRPCClient struct {
    client pbAuth.AuthServiceClient
}

// Auth client
func NewAuthServiceGRPCClient(conn *grpc.ClientConn) AuthServiceClient {
    return &authServiceGRPCClient{
        client: pbAuth.NewAuthServiceClient(conn),
    }
}

// Auth client- authenicate
func (c *authServiceGRPCClient) Authenticate(ctx context.Context, req interface{}) (interface{}, error) {
    authReq, ok := req.(*pbAuth.AuthenticateRequest)
    if !ok {
        return nil, status.Errorf(codes.Internal, "Failed to assert request type")
    }

    resp, err := c.client.Authenticate(ctx, authReq)
    if err != nil {// Handle gRPC error
        return nil, status.Errorf(codes.Internal, "Failed to authenticate: %v", err)
    }
    return resp, nil
}
