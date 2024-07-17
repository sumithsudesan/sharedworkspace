package metrics

import (
	"encoding/json"
    "log"
    "net/http"
    "os"
	"strings"
)

// Server - server
type Server struct {
	svrConfig   Config
    dbClient    *Database
}

// NewServer create an instance of metrics.Server object
func NewServer( config Config) (*Server, error) {
	server := Server{
		svrConfig: config,
        dbClient: NewInfluxDBClient(config),
	}
	return &server, nil
}

// Start API will starts the metrics server
func (svr *Server) Start() {

    // HTTP
    // Initialize HTTP server
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/v1/metrics", svr.handlePlaceOrder).Methods("POST")

    // Setup CORS
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", 
										 "POST",
										 "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(router)

    log.Println("[INFO] Starting metrics collector Service")
    log.Fatal(http.ListenAndServe(svrConfig.OrderServiceHTTPPort(), corsHandler))
}

// Save order handler
func (svr *Server) handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
		var m Metric
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := svr.dbClient.WriteMetric(r.Context(), m); err != nil {
			http.Error(w, "Failed to write the DB", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// jsonResponse utility function to send JSON response
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}