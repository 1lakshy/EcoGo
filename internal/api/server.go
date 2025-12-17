package api

import (
	"fmt"
	"go-ecommerce-api/configs"
	"go-ecommerce-api/internal/api/rest"
	"go-ecommerce-api/internal/api/rest/handler"
	"go-ecommerce-api/internal/domain"
	"go-ecommerce-api/internal/helper"
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

	// run auto migrate
	db.AutoMigrate(&domain.User{})
	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:app,
		DB: db,
		Auth: auth,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)
}


func setupRoutes( rh * rest.RestHandler){
	// user handler
	handler.SetupUserRoutes(rh )
}