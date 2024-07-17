package metrics

import (
    "context"
    "time"

    "github.com/influxdata/influxdb-client-go/v2"
)

// InfluxDB
// Based on Database interface
type InfluxDBClient struct {
    client   influxdb2.Client
    writeAPI influxdb2.WriteAPIBlocking
}

// Create new instance InfluxDBClient
func NewInfluxDBClient(config *Config) *InfluxDBClient {
	client := influxdb2.NewClient(config.InfluxDBURL(), config.InfluxDBToken())
	writeAPI := client.WriteAPIBlocking(config.InfluxDBOrg(), config.InfluxDBBucket())

	return &InfluxDBClient{
		client:   client,
		writeAPI: writeAPI,
	}
}

// write metrivs to influxdb
// Database interface
func (c *InfluxDBClient) WriteMetric(ctx context.Context, metric metrics.Metric) error {
    p := influxdb2.NewPoint(
        metric.Metric,
        map[string]string{"service": metric.Service},
        map[string]interface{}{"value": metric.Value},
        time.Now(),
    )
    return c.writeAPI.WritePoint(ctx, p)
}