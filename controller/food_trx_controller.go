package controller

import (
	"github.com/fikryfahrezy/ffood/exception"
	"github.com/fikryfahrezy/ffood/middleware"
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/service"
	"github.com/gofiber/fiber/v2"
)

type FoodTrxController struct {
	FoodTrxService service.FoodTrxService
}

func (Controller FoodTrxController) InsertFoodTrx(c *fiber.Ctx) error {
	request := new(model.InsertFoodTrxRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}

	response, err := Controller.FoodTrxService.InsertFoodTrx(model.InsertFoodTrxRequest{
		FoodId: request.FoodId,
	}, int64(c.Locals("id").(uint64)))
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

func (Controller FoodTrxController) Route(App fiber.Router) {
	router := App.Group("/foodtransactions")
	router.Post("", middleware.CheckToken(), Controller.InsertFoodTrx)
}

func NewFoodTrxController(FoodTrxService *service.FoodTrxService) FoodTrxController {
	return FoodTrxController{
		FoodTrxService: *FoodTrxService,
	}
}
