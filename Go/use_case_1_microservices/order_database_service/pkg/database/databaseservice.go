package database

import (
    "context"
    "log"
    "net"
    "time"

    "google.golang.org/grpc"
    pb "order_database_service/data-definitions/order"
)

// Database service
type DatabaseService struct {
    pb.UnimplementedOrderDatabaseServiceServer
    db db.Database
}

// Database service
func NewDatabaseService(cfg *config.Config) (*DatabaseService, error) {
    var database db.Database
    var err error

    switch cfg.DatabaseType() {
    case "mongodb":
        database, err = db.NewMongoDB(cfg.MongoDBURI(), cfg.MongoDBName(), time.Duration(cfg.QueryTimeoutSeconds())*time.Second)
    case "postgres":
        database, err = db.NewPostgreSQL(cfg.PostgreSQLDSN())
    default:
        return nil, fmt.Errorf("Invalid database type: %s", cfg.DatabaseType())
    }
    if err != nil {
        return nil, err
    }

    return &Server{db: database}, nil
}

// Save the order
func (s *DatabaseService) SaveOrder(ctx context.Context, req *pb.SaveOrderRequest) (*pb.SaveOrderResponse, error) {
    return s.db.SaveOrder(ctx, req)
}

// Get the order
func (s *DatabaseService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
    return s.db.GetOrder(ctx, req.OrderId)
}

// Delete Order
func (s *DatabaseService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
    return s.db.DeleteOrder(ctx, req.OrderId)
}