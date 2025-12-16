package handler

import (
	"go-ecommerce-api/internal/api/rest"
	"go-ecommerce-api/internal/dto"
	"go-ecommerce-api/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"

)

type UserHandler struct {
	svc service.UserService
	
}


func SetupUserRoutes(rh * rest.RestHandler){
	app := rh.App


	svc := service.UserService{	}
	handler := UserHandler{
		svc: svc ,
	}
	// public Endpoints
    app.Post("/register",handler.Register)

	app.Post("/login", handler.Login)

	// private endpoint
	app.Get("/verify",handler.GetVerificationCode)
	app.Post("/verify",handler.Verify)
	app.Post("/profile",handler.CreateProfile)
	app.Get("/profile",handler.GetProfile)

	app.Post("/cart",handler.AddToCart)
	app.Get("/cart",handler.GetCart)
    app.Get("/order",handler.GetOrder)
	app.Get("/order/:id", handler.GetOrder)


}


func (h * UserHandler) Register(ctx *fiber.Ctx) error{
	
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)

	if err != nil{
    return ctx.Status(http.StatusBadGateway).JSON(&fiber.Map{
		"message":"please provide a valid input",
	})
	}

	token,err := h.svc.SignUp(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"error in signup",
		})
	}
	
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Login",
	}) 
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"GetVerificationCode",
	}) 
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Verify",
	}) 
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"CreateProfile",
	}) 
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"GetProfile",
	}) 
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"CreateProfile",
	}) 
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"Get cart",
	}) 
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"CreateProfile",
	}) 
}

// func (h *UserHandler) GetCart(ctx *fiber.Ctx) error{
// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message":"CreateProfile",
// 	}) 
// }

func (h *UserHandler) CreateOrders(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"CreateProfile",
	}) 
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"CreateProfile",
	}) 
}