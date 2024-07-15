package main

import (
	"auth_service/pkg/auth"
	"fmt"
)

func init() {
	InitConfig()
}

func main() {
	var authsvr *auth.Server = nil	
	config:=auth.LoadConfig()
	// grpc server
	authsvr, err = auth.NewServer(
		config,
	)
	if err != nil {
		fmt.Println("[ERROR] Exiting server, error : %v", err)
		return
	}
	authsvr.Start()
}
