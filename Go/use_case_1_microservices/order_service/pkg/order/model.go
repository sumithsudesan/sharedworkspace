package order

// OrderRequest - represents the incoming order request
type OrderRequest struct {
    UserID    string  `json:"user_id"`
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

// OrderResponse - represents the response
type OrderResponse struct {
    OrderID string `json:"order_id"`
    Message string `json:"message"`
}

// OrderDatabaseServiceClient defines the interface for interacting with the order database service.
type OrderDatabaseServiceClient interface {
    PlaceOrder(ctx context.Context, req interface{}) (interface{}, error)
    GetOrder(ctx context.Context, req interface{}) (interface{}, error)
    DeleteOrder(ctx context.Context, req interface{}) (interface{}, error)
}

// AuthServiceClient defines the interface for interacting with the auth service.
type AuthServiceClient interface {
    Authenticate(ctx context.Context, req interface{}) (interface{}, error)
}