package database

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

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(timeoutInterceptor))
	
	pb.RegisterOrderDatabaseServiceServer(grpcServer, &DatabaseService{svr.svrConfig})

	// Start the gRPC server in a separate goroutine
	go func() {
		log.Printf("[INFO] Starting auth data gRPC server on - %s...", svr.svrConfig.GRPCPort())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("[ERROR] Failed to serve auth data gRPC server: %v", err)
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

//
func timeoutInterceptor(    ctx context.Context,
							req interface{},
							info *grpc.UnaryServerInfo,
							handler grpc.UnaryHandler,) (interface{}, error) {
    // Set a timeout for the request
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()
    return handler(ctx, req)
}