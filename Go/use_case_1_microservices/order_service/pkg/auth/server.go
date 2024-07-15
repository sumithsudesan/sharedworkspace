package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Server -  server
type Server struct {
	svrConfig       Config
}

// NewServer create an instance of auth.Server object
func NewServer( config Config) (*AuthServer, error) {
	server := AuthServer{
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

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register the HandleAuth server implementation with the gRPC server
	src.RegisterAuthServiceServer(grpcServer, &auth.AuthService{})

	// Start the gRPC server in a separate goroutine
	go func() {
		log.Printf("[INFO] Starting gRPC server on - %s...", svr.svrConfig.GRPCPort())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("[ERROR] Failed to serve gRPC server: %v", err)
		}
	}()

	// Trap signals (Ctrl+C) to gracefully shutdown the server
	// This is optional but recommended for graceful shutdowns
	// For example, handle os.Interrupt or syscall.SIGTERM

	// Block the main goroutine until a signal is received
	// This prevents the program from exiting immediately
	// after starting the server
	<-make(chan struct{})

	// Gracefully shutdown the gRPC server
	grpcServer.GracefulStop()
}
