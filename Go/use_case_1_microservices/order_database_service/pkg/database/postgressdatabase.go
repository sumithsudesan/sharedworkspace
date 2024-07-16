package database

import (
    "context"
    "database/sql"
    _ "github.com/lib/pq"
    "order_database_service/data-definitions/order"
)

// Postgres SQL
type PostgreSQL struct {
    db *sql.DB
}

// 
func NewPostgreSQL(dsn string) (*PostgreSQL, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    return &PostgreSQL{db: db}, nil
}

// Save the order
func (p *PostgreSQL) SaveOrder(ctx context.Context, order *proto.SaveOrderRequest) (*proto.SaveOrderResponse, error) {
    var orderID string
    query := `INSERT INTO orders (user_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING order_id`
    err := p.db.QueryRowContext(ctx, query, order.UserId, order.ProductId, order.Quantity, order.Price).Scan(&orderID)
    if err != nil {
        return nil, err
    }
    return &proto.SaveOrderResponse{OrderId: orderID, Message: "Order saved"}, nil
}

// Get order details
func (p *PostgreSQL) GetOrder(ctx context.Context, orderID string) (*proto.GetOrderResponse, error) {
    var order proto.GetOrderResponse
    query := `SELECT order_id, user_id, product_id, quantity, price FROM orders WHERE order_id = $1`
    row := p.db.QueryRowContext(ctx, query, orderID)
    err := row.Scan(&order.OrderId, &order.UserId, &order.ProductId, &order.Quantity, &order.Price)
    if err != nil {
        return nil, err
    }
    return &order, nil
}

// Delete order
func (p *PostgreSQL) DeleteOrder(ctx context.Context, orderID string) (*proto.DeleteOrderResponse, error) {
    query := `DELETE FROM orders WHERE order_id = $1`
    _, err := p.db.ExecContext(ctx, query, orderID)
    if err != nil {
        return nil, err
    }
    return &proto.DeleteOrderResponse{Message: "Order deleted"}, nil
}