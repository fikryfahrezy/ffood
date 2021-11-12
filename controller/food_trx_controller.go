package controller

import (
	"github.com/fikryfahrezy/ffood/service"
	"github.com/gofiber/fiber/v2"
)

type FoodTrxController struct {
	FoodTrxService service.FoodTrxService
}

func NewFoodTrxController(FoodTrxService *service.FoodTrxService) FoodTrxController {
	return FoodTrxController{
		FoodTrxService: *FoodTrxService,
	}
}

func (Controller FoodTrxController) Route(App fiber.Router) {
	App.Group("/foodtransactions")
}
