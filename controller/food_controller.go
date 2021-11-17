package controller

import (
	"github.com/fikryfahrezy/ffood/exception"
	"github.com/fikryfahrezy/ffood/middleware"
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/service"

	"github.com/gofiber/fiber/v2"
)

type FoodController struct {
	FoodService service.FoodService
}

func (Controller FoodController) InsertFood(c *fiber.Ctx) error {
	request := new(model.InsertFoodRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}
	response, err := Controller.FoodService.InsertFood(model.InsertFoodRequest{
		Name: request.Name,
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

func (Controller FoodController) GetFoods(c *fiber.Ctx) error {
	response, err := Controller.FoodService.GetAllFood()
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

func (Controller FoodController) GetFood(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := Controller.FoodService.GetFood(id)
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

func (Controller FoodController) DeleteFood(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := Controller.FoodService.DeleteFood(id, int64(c.Locals("id").(uint64)))
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

func (Controller FoodController) UpdateFood(c *fiber.Ctx) error {
	id := c.Params("id")
	request := new(model.UpdateFoodRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.ErrorHandler(c, err)
	}
	response, err := Controller.FoodService.UpdateFood(model.UpdateFoodRequest{
		Name: request.Name,
	}, id, int64(c.Locals("id").(uint64)))
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

func (Controller FoodController) Route(App fiber.Router) {
	router := App.Group("/foods")
	router.Post("", middleware.CheckToken(), middleware.CheckRole("seller"), Controller.InsertFood)
	router.Get("", Controller.GetFoods)
	router.Get("/:id", Controller.GetFood)
	router.Delete("/:id", middleware.CheckToken(), middleware.CheckRole("seller"), Controller.DeleteFood)
	router.Patch("/:id", middleware.CheckToken(), middleware.CheckRole("seller"), Controller.UpdateFood)
}

func NewFoodController(FoodService *service.FoodService) FoodController {
	return FoodController{
		FoodService: *FoodService,
	}
}
