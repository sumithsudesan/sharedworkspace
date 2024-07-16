package database

import (
    "context"
    "order_database_service/data-definitions/order"
)
// Interface for DB 
type Database interface {
    SaveOrder(ctx context.Context, order *proto.SaveOrderRequest) (*proto.SaveOrderResponse, error)
    GetOrder(ctx context.Context, orderID string) (*proto.GetOrderResponse, error)
    DeleteOrder(ctx context.Context, orderID string) (*proto.DeleteOrderResponse, error)
}