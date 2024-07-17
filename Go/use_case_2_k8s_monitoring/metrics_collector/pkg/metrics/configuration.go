package metrics

import (
	"log"
	"os"
)

type Config struct {
	port            string
	influxDBURL     string
	influxDBOrg     string
	influxDBBucket  string
	influxDBToken   string
}

func NewConfig() *Config {
	return &Config{
		port:           getEnv("METRICS_COLLECTOR_PORT", "8080"),
		influxDBURL:    getEnv("INFLUXDB_URL", "http://localhost:8086"),
		influxDBOrg:    getEnv("INFLUXDB_ORG", "my-org"),
		influxDBBucket: getEnv("INFLUXDB_BUCKET", "metrics"),
		influxDBToken:  getEnv("INFLUXDB_TOKEN", "###rrrrr"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Get Service port
func (c *Config) Port() string {
	return c.port
}

// Get influxDB URL
func (c *Config) InfluxDBURL() string {
	return c.influxDBURL
}

// Get InfluxDB Org
func (c *Config) InfluxDBOrg() string {
	return c.influxDBOrg
}

// Get influx DB Bucket
func (c *Config) InfluxDBBucket() string {
	return c.influxDBBucket
}

// Get influxdb token
func (c *Config) InfluxDBToken() string {
	return c.influxDBToken
}