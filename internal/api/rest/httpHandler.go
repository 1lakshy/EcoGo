package rest

import (
	"go-ecommerce-api/internal/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type RestHandler struct {
	App *fiber.App
	DB *gorm.DB 
	Auth helper.Auth
}