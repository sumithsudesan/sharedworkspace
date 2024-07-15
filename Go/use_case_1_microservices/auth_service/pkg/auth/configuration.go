package auth

import (
    "log"

    "github.com/spf13/viper"
)

type Config struct {
    gRPCPort     string
    mySQLDSN     string
    jWTSecretKey string
    authDataServiceEndpoint string
}

func LoadConfig() Config {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/etc/auth_service/")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("[ERROR] Error reading config file, %s", err)
    }

    var config Config
    err := viper.Unmarshal(&config)
    if err != nil {
        log.Fatalf("[ERROR] Unable to decode into struct, %v", err)
    }

    return config
}

// Get GRPC Port 
func (c *Config) GRPCPort() string {
	return c.gRPCPort
}

// Get MySQL DSN
func (c *Config) MySQLDSN() string {
	return c.mySQLDSN
}

// JWT Secret Key
func (c *Config) JWTSecretKey() string {
	return c.jWTSecretKey
}

// Auth DataService Endpoint
func (c *Config) AuthDataServiceEndpoint() string {
	return c.authDataServiceEndpoint
}