package main

import (
	"go-ecommerce-api/configs"
	"go-ecommerce-api/internal/api"
)

func main() {

	cfg ,err := configs.SetupEnv()

	if err != nil {
		panic(err)
	}

    api.StartServer(cfg)
}