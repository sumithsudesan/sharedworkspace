package database

import (
    "log"

    "github.com/spf13/viper"
)


// Config for database
type Config struct {
    dbHost         string
    dbport         string
    dbUser         string
    dbPassword     string
    dbName         string
    queryTimeoutMS int 
    driver         string // "mysql" or "postgres"
    GRPCPort       string
}

// loads db configuration
func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/etc/auth-database/")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}

// DBHost
func (c *Config) DBHost() string {
	return c.dbHost
}

// DBPort
func (c *Config) DBPort() string {
	return c.dbport
}

// DBUser
func (c *Config) DBUser() string {
	return c.dbUser
}

// DBPassword
func (c *Config) DBPassword() string {
	return c.dbPassword
}

// DBName
func (c *Config) DBName() string {
	return c.dbName
}

// QueryTimeoutMS
func (c *Config) QueryTimeoutMS() int {
	return c.queryTimeoutMS
}

// Driver
func (c *Config) Driver() string {
	return c.driver
}

// Get GRPC Port 
func (c *Config) GRPCPort() string {
	return c.GRPCPort
}
