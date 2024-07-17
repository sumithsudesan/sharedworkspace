package Metrics

import (
    "context"
)

// metrics 
type Metric struct {
    Service string  `json:"service"`
    Metric  string  `json:"metric"`
    Value   float64 `json:"value"`
}

// Interface for database
type Database interface {
    WriteMetric(ctx context.Context, metric metrics.Metric) error
}