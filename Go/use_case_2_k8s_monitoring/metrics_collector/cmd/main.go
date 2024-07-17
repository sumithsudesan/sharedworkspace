package main

import (
	"metrics_collector/pkg/metrics"
	"fmt"
)

func init() {
	InitConfig()
}

func main() {
	var metricsSvr *metrics.Server = nil	
	config:=order.LoadConfig()
	// metricsSvr server
	metricsSvr, err = order.NewServer(
		config,
	)
	if err != nil {
		fmt.Println("[ERROR] Exiting Metrics server, error : %v", err)
		return
	}
	metricsSvr.Start()
}
