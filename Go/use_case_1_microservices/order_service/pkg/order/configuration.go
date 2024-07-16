package order

import (
    "os"
)


// Config for order service
type Config struct {
    authServicePort       string
    authServiceName       string
    orderDatabasePort     string
    orderDatabaseService  string
    orderServiceHTTPPort  string
}

// 
func LoadConfig() Config {
    config := Config{
        authServicePort:      getEnv("AUTH_SERVICE_PORT", "50051"),
        authServiceName:      getEnv("AUTH_SERVICE_NAME", "auth-service"),
        orderDatabasePort:    getEnv("ORDER_DATABASE_PORT", "50052"),
        orderDatabaseService: getEnv("ORDER_DATABASE_SERVICE", "order-database-service"),
        orderServiceHTTPPort: getEnv("ORDER_SERVICE_HTTP_PORT", "8080"),
    }
    return config
}

// Get enviornment variable
func getEnv(key, defVal string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return defVal
}

// Get Auth Service Port
func (c *Config) AuthServicePort() string {
	return c.authServicePort
}

// Get Auth Service 
func (c *Config) AuthServiceName() string {
	return c.authServiceName
}

// Get data Service port
func (c *Config) OrderDatabasePort() string {
	return c.orderDatabasePort
}

// Get data Service 
func (c *Config) OrderDatabaseService() string {
	return c.orderDatabaseService
}

// Get order Service port
func (c *Config) OrderServiceHTTPPort() string {
	return c.orderServiceHTTPPort
}