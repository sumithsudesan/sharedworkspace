package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "order_database_service/data-definitions/order"
)

// MongoDB
type MongoDB struct {
    client   *mongo.Client
    database *mongo.Database
    timeout  time.Duration
}

// 
func NewMongoDB(uri, dbName string, timeout time.Duration) (*MongoDB, error) {
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        return nil, err
    }
    return &MongoDB{client: client, database: client.Database(dbName), timeout: timeout}, nil
}

// Save order
func (m *MongoDB) SaveOrder(ctx context.Context, order *proto.SaveOrderRequest) (*proto.SaveOrderResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, m.timeout)
    defer cancel()
    collection := m.database.Collection("orders")
    res, err := collection.InsertOne(ctx, bson.M{
        "user_id":    order.UserId,
        "product_id": order.ProductId,
        "quantity":   order.Quantity,
        "price":      order.Price,
    })
    if err != nil {
        return nil, err
    }
    return &proto.SaveOrderResponse{OrderId: res.InsertedID.(string), Message: "Order saved"}, nil
}

// Get order
func (m *MongoDB) GetOrder(ctx context.Context, orderID string) (*proto.GetOrderResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, m.timeout)
    defer cancel()
    collection := m.database.Collection("orders")
    var result bson.M
    err := collection.FindOne(ctx, bson.M{"_id": orderID}).Decode(&result)
    if err != nil {
        return nil, err
    }
    return &proto.GetOrderResponse{
        OrderId:   result["_id"].(string),
        UserId:    result["user_id"].(string),
        ProductId: result["product_id"].(string),
        Quantity:  result["quantity"].(int32),
        Price:     result["price"].(float64),
    }, nil
}

// Delete order
func (m *MongoDB) DeleteOrder(ctx context.Context, orderID string) (*proto.DeleteOrderResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, m.timeout)
    defer cancel()
    collection := m.database.Collection("orders")
    _, err := collection.DeleteOne(ctx, bson.M{"_id": orderID})
    if err != nil {
        return nil, err
    }
    return &proto.DeleteOrderResponse{Message: "Order deleted"}, nil
}