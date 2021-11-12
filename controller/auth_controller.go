package controller

import (
	"github.com/fikryfahrezy/ffood/exception"
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(AuthService *service.AuthService) AuthController {
	return AuthController{
		AuthService: *AuthService,
	}
}

func (Controller AuthController) Route(App fiber.Router) {
	router := App.Group("/auth")
	router.Post("/register", Controller.Register)
	router.Post("/login", Controller.Login)
}

func (Controller AuthController) Register(c *fiber.Ctx) error {
	request := new(model.RegisterRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := Controller.AuthService.Register(model.RegisterRequest{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Error:  nil,
	})
}

func (Controller AuthController) Login(c *fiber.Ctx) error {
	request := new(model.AuthRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, verified, err := Controller.AuthService.Login(model.AuthRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return exception.ErrorHandler(c, err)
	}
	if !verified {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   nil,
			Error: model.GeneralError{
				General: "AUTHENTICATION_FAILURE",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Error:  nil,
	})
}
