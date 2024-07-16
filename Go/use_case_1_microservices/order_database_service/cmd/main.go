package main

import (
	"auth_service/pkg/database"
	"fmt"
)

func init() {
	InitConfig()
}

func main() {
	var dbSvr *database.Server = nil	
	config:=auth.LoadConfig()
	// grpc server
	dbSvr, err = database.NewServer(
		config,
	)
	if err != nil {
		fmt.Println("[ERROR] Exiting server, error : %v", err)
		return
	}
	dbSvr.Start()
}
