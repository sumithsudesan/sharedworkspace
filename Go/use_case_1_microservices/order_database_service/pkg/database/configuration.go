package database

import (
    "log"
    "time"
    "github.com/spf13/viper"
)

type Config struct {
    gRPCPort             string
    databaseType         string
    mongoDBURI           string
    mongoDBName          string
    postgreSQLDSN        string
    queryTimeoutSeconds  int
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.AddConfigPath("/configs")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
        return nil, err
    }

    config := &Config{
        GRPCPort:            viper.GetString("grpc_port"),
        DatabaseType:        viper.GetString("database_type"),
        MongoDBURI:          viper.GetString("mongodb_uri"),
        MongoDBName:         viper.GetString("mongodb_name"),
        PostgreSQLDSN:       viper.GetString("postgresql_dsn"),
        QueryTimeoutSeconds: viper.GetInt("query_timeout_seconds"),
    }

    return config, nil
}

// Get GRPC Port 
func (c *Config) GRPCPort() string {
	return c.gRPCPort
}

// Get DB type
func (c *Config) DatabaseType() string {
	return c.databaseType
}

// Get mongo DB URI
func (c *Config) MongoDBURI() string {
	return c.mongoDBURI
}

// Get Mongo DB Name
func (c *Config) MongoDBName() string {
	return c.mongoDBName
}

// Get postgre SQL DSN
func (c *Config) PostgreSQLDSN() string {
	return c.postgreSQLDSN
}

// Get Query Timeout Seconds
func (c *Config) QueryTimeoutSeconds() int {
	return c.queryTimeoutSeconds
}