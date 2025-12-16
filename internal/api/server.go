package api

import (
	"fmt"
	"go-ecommerce-api/configs"
	"go-ecommerce-api/internal/api/rest"
	"go-ecommerce-api/internal/api/rest/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config configs.AppConfig) {

	fmt.Println("Hello, World!")

	log.Printf("config %v",config.Dsn )

	db,err := gorm.Open(postgres.Open(config.Dsn),&gorm.Config{})

	if err!= nil{
		log.Fatal("Database Connection error %v\n",err)
	}

	log.Println("Database Connected")
	log.Print(db)
	app := fiber.New()

	rh := &rest.RestHandler{
		App:app,
		DB: db,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)
}


func setupRoutes( rh * rest.RestHandler){
	// user handler
	handler.SetupUserRoutes(rh )
}