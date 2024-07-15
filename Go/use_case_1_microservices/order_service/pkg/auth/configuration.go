package auth

import (
    "log"

    "github.com/spf13/viper"
)

type Config struct {
    GRPCPort     string
    MySQLDSN     string
    JWTSecretKey string
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
	return c.GRPCPort
}

// Get MySQL DSN
func (c *Config) MySQLDSN() string {
	return c.MySQLDSN
}

// JWT Secret Key
func (c *Config) JWTSecretKey() string {
	return c.JWTSecretKey
}