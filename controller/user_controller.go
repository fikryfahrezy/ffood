package controller

import (
	"golang-simple-boilerplate/exception"
	"golang-simple-boilerplate/middleware"
	"golang-simple-boilerplate/model"
	"golang-simple-boilerplate/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{
		UserService: *UserService,
	}
}

func (Controller UserController) Route(App fiber.Router) {
	router := App.Group("/user")
	router.Get("/profile", middleware.CheckToken(), Controller.Profile)
	router.Patch("/profile", middleware.CheckToken(), Controller.UpdateProfile)
}

func (Controller UserController) Profile(c *fiber.Ctx) error {
	response, err := Controller.UserService.Profile(model.ProfileRequest{
		Email: c.Locals("email").(string),
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

func (Controller UserController) UpdateProfile(c *fiber.Ctx) error {
	request := new(model.UpdateProfileRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := Controller.UserService.UpdateProfile(model.UpdateProfileRequest{
		Email:    c.Locals("email").(string),
		Name:     request.Name,
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
