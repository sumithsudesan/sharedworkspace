package main

import (
	"auth_service/pkg/order"
	"fmt"
)

func init() {
	InitConfig()
}

func main() {
	var orderSvr *order.Server = nil	
	config:=order.LoadConfig()
	// grpc server
	orderSvr, err = order.NewServer(
		config,
	)
	if err != nil {
		fmt.Println("[ERROR] Exiting Order server, error : %v", err)
		return
	}
	orderSvr.Start()
}
