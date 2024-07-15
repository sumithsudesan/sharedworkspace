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
	// initlaise DB connection based on the config
	db, err := svr.InitializeDatabase() 
    if err != nil {
        log.Fatalf("[ERROR] Failed to initialize database: %v", err)
    }

    // Initialize UserRepository (MySQL)
	// can add  new database type by the configuration
	//
    userRepository := database.NewMySQLUserRepository(db)

	// GRPC handler 
	server := database.NewAuthDatabaseServiceServer(userRepository)

	// Register the HandleAuth server implementation with the gRPC server
	src.RegisterAuthServiceServer(grpcServer, server)

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
func (svr *Server) InitializeDatabase() (*sql.DB, error) {
	var dsn string
	if config.Driver() == "mysql" {
            // Create MySQL DSN 
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
								svr.svrConfig.DBUser(),
								svr.svrConfig.DBPassword(),
								svr.svrConfig.DBHost(),
								svr.svrConfig.DBPort(),
								svr.svrConfig.DBName(),	)
    	}
    } else {
        return nil, fmt.Errorf("Unsupported database driver: %s", config.config.Driver())
    }

    // Open database 	
	db, err := sql.Open(config.Driver(), dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to open database connection: %v", err)
    }

	db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    // verify connection
    if err := db.Ping(); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    return db, nil
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