package order

import (
	"encoding/json"
    "log"
    "net/http"
    "os"
	"strings"

    "github.com/gorilla/mux"
    "github.com/dgrijalva/jwt-go"
)

// Server - server
type Server struct {
	svrConfig   Config
    authClient AuthServiceClient
    orderDBClient OrderDatabaseServiceClient
}

// NewServer create an instance of order.Server object
func NewServer( config Config) (*Server, error) {
	server := Server{
		svrConfig: config
	}
	return &server, nil
}

// Start API will starts the order server
func (svr *Server) Start() {
    // for auth service 
    // Initialize gRPC client connection to auth_service
    authConn, err := grpc.Dial(fmt.Sprintf("%s:%s", 
                                            svr.svrConfig.AuthServiceName(),  
                                                svr.svrConfig.AuthServicePort())
                                , grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial auth_service: %v", err)
    }

    defer authConn.Close()

    // Create gRPC clients using interface implementations
    svr.authClient := NewAuthServiceGRPCClient(authConn)

    // for db service
    // Initialize gRPC client connection to order_database_service
    orderDBConn, err := grpc.Dial(fmt.Sprintf("%s:%s", 
                                                svr.svrConfig.OrderDatabaseService(),  
                                                    svr.svrConfig.OrderDatabasePort()),
                                grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial order_database_service: %v", err)
    }
    defer orderDBConn.Close()

    // Create gRPC clients using interface implementations
    svr.orderDBClient := NewOrderDatabaseGRPCClient(orderDBConn)


    // HTTP
    // Initialize HTTP server
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/order", svr.handlePlaceOrder).Methods("POST")
    router.HandleFunc("/order/{orderId}", svr.handleListOrder).Methods("GET")
    router.HandleFunc("/order/{orderId}",  svr.handleDeleteOrder).Methods("DELETE")

    // Setup CORS
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", 
										 "POST",
										 "PUT",
										 "DELETE", 
										 "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(router)

    log.Println("[INFO] Starting Order Service")
    log.Fatal(http.ListenAndServe(svrConfig.OrderServiceHTTPPort(), corsHandler))
}

// Save order handler
func (svr *Server) handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse JSON request body
        var req pb.PlaceOrderRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        // Validate JWT token via auth_service
        valid, err := validateToken( &pbAuth.AuthenticateRequest{Token: req.Token})
        if err != nil || !valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Call gRPC method to save order
        resp, err := svr.orderDBClient.PlaceOrder(context.Background(), &req)
        if err != nil {
            http.Error(w, "Failed to save order", http.StatusInternalServerError)
            return
        }

        // Respond with success
        jsonResponse(w, http.StatusOK, resp)
    }
}

// List order handler
func (svr *Server) handleListOrder(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract orderId from URL
        orderId := mux.Vars(r)["orderId"]
        token := r.Header.Get("Authorization")

        // Validate JWT token via auth_service
        valid, err := validateToken(&pbAuth.AuthenticateRequest{Token: token})
        if err != nil || !valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Call gRPC method to get order
        resp, err := svr.orderDBClient.GetOrder(context.Background(), &pb.GetOrderRequest{
            Token:   token,
            OrderId: orderId,
        })
        if err != nil {
            http.Error(w, "Failed to get order", http.StatusInternalServerError)
            return
        }

        // Respond with order details
        jsonResponse(w, http.StatusOK, resp)
    }
}

// Delete order handler
func (svr *Server) handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract orderId from URL
        orderId := mux.Vars(r)["orderId"]
        token := r.Header.Get("Authorization")

        // Validate JWT token via auth_service
        valid, err := validateToken( &pbAuth.AuthenticateRequest{Token: token})
        if err != nil || !valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Call gRPC method to delete order
        resp, err := svr.orderDBClient.DeleteOrder(context.Background(), &pb.DeleteOrderRequest{
            Token:   token,
            OrderId: orderId,
        })
        if err != nil {
            http.Error(w, "Failed to delete order", http.StatusInternalServerError)
            return
        }

        // Respond with success
        jsonResponse(w, http.StatusOK, resp)
    }
}

// validateToken validates JWT token via auth_service
func (svr *Server) validateToken( req interface{}) (bool, error) {
    // Assert the request to the expected type
    authReq := req.(*pbAuth.AuthenticateRequest)
    resp, err := svr.authClient.Authenticate(context.Background(), authReq)
    if err != nil {
        return false, err
    }
    return resp.Valid, nil
}

// jsonResponse utility function to send JSON response
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}