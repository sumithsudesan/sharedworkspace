package order

import (
    "context"

    "google.golang.org/grpc"
    pb "data-definitions/order"
)

// Order DB client
type orderDatabaseGRPCClient struct {
    client pb.OrderDatabaseServiceClient
}

func NewOrderDatabaseGRPCClient(conn *grpc.ClientConn) OrderDatabaseServiceClient {
    return &orderDatabaseGRPCClient{
        client: pb.NewOrderDatabaseServiceClient(conn),
    }
}

// Order DB client- Place
func (c *orderDatabaseGRPCClient) PlaceOrder(ctx context.Context, req interface{}) (interface{}, error) {
    placeOrderReq, ok := req.(*pb.PlaceOrderRequest)
    if !ok {
        return nil, status.Errorf(codes.Internal, "Failed to assert request type")
    }

    resp, err := c.client.PlaceOrder(ctx, placeOrderReq)
    if err != nil {// Handle gRPC call error
        return nil, status.Errorf(codes.Internal, "Failed to place order: %v", err)
    }
    return resp, nil
}

// Order DB client- Get
func (c *orderDatabaseGRPCClient) GetOrder(ctx context.Context, req interface{}) (interface{}, error) {
    getOrderReq, ok := req.(*pb.GetOrderRequest)
    if !ok {
        return nil, status.Errorf(codes.Internal, "Failed to assert request type")
    }

    resp, err := c.client.GetOrder(ctx, getOrderReq)
    if err != nil {// Handle gRPC call error
        
        return nil, status.Errorf(codes.Internal, "Failed to get order: %v", err)
    }
    return resp, nil
}

// Order DB client- Delete
func (c *orderDatabaseGRPCClient) DeleteOrder(ctx context.Context, req interface{}) (interface{}, error) {
    deleteOrderReq, ok := req.(*pb.DeleteOrderRequest)
    if !ok {
        return nil, status.Errorf(codes.Internal, "Failed to assert request type")
    }

    resp, err := c.client.DeleteOrder(ctx, deleteOrderReq)
    if err != nil {// Handle gRPC call error        
        return nil, status.Errorf(codes.Internal, "Failed to delete order: %v", err)
    }
    return resp, nil
}
