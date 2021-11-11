package controller

import (
	"golang-simple-boilerplate/middleware"
	"golang-simple-boilerplate/service"

	"github.com/gofiber/fiber/v2"
)

type FoodController struct {
	FoodService service.FoodService
}

func NewFoodController(FoodService *service.FoodService) FoodController {
	return FoodController{
		FoodService: *FoodService,
	}
}

func (Controller FoodController) Route(App fiber.Router) {
	router := App.Group("/foods")
	router.Get("/profile", middleware.CheckToken(), func(ctx *fiber.Ctx) error {
		err := ctx.SendString("hi")
		return err
	})
}
