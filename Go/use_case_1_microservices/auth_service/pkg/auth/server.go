package auth

import (
    "context"
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

   	pb "data-definitions/auth"
	pc "data-definitions/user"
)
// Server -  server
type Server struct {
	svrConfig       Config
}

// NewServer create an instance of auth.Server object
func NewServer( config Config) (*Server, error) {
	server := Server{
		svrConfig: config
	}
	return &server, nil
}

// Start API will starts the AuthServer
func (svr *Server) Start() {
	// Create a listener for the specified GRPC port
	listener, err := net.Listen("tcp", svr.svrConfig.GRPCPort())
	if err != nil {
		log.Fatalf("[ERROR] Failed to listen on port %s, error: %v", svr.svrConfig.GRPCPort(), err)
	}
	defer listener.Close()

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Initialize gRPC client for AuthDataService
	if client, err := svr.InitializeGRPCClient() 	
	if err != nil {
        log.Fatalf("[ERROR] Failed to initialize gRPC client: %v", err)
    }

	// Register the AuthService server implementation with the gRPC server
	auth.RegisterAuthServiceServer(grpcServer, auth.NewAuthService(svr.svrConfig, client))

	// Start the gRPC server in a separate goroutine
	go func() {
		log.Printf("[INFO] Starting gRPC server on - %s...", svr.svrConfig.GRPCPort())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("[ERROR] Failed to serve gRPC server: %v", err)
		}
	}()

    // Handle graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    log.Println("[INFO] Shutting down gRPC server...")
    grpcServer.GracefulStop()
    log.Println("[INFO] gRPC server stopped.")
}


func (svr *Server) InitializeGRPCClient() (pb.AuthDataServiceClient,error) {
    // Establish connection to the AuthDataService
    conn, err := grpc.Dial(svr.svrConfig.AuthDataServiceEndpoint(), grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

 	// Create a gRPC client for AuthDataService
    return pc.NewAuthDataServiceClient(conn), nil
}

// RegisterAuthServiceServer with AuthService .
func RegisterAuthServiceServer(s *grpc.Server, service *AuthService) {
    pb.RegisterAuthServiceServer(s, service)
}