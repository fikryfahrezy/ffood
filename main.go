package main

import (
	"os"

	"github.com/fikryfahrezy/ffood/config"
	"github.com/fikryfahrezy/ffood/controller"
	"github.com/fikryfahrezy/ffood/exception"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	mysql := config.MysqlConnection()

	authRepository := repository.NewAuthRepository(mysql)
	userRepository := repository.NewUserRepository(mysql)
	foodRepository := repository.NewFoodRepository(mysql)
	foodTrxRepository := repository.NewFoodTrxRepository(mysql)

	authService := service.NewAuthService(&authRepository)
	userService := service.NewUserService(&userRepository)
	foodService := service.NewFoodService(&foodRepository)
	foodTrxService := service.NewFoodTrxService(&foodTrxRepository)

	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)
	foodController := controller.NewFoodController(&foodService)
	foodTrxController := controller.NewFoodTrxController(&foodTrxService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/api/v1")
	authController.Route(v1)
	userController.Route(v1)
	foodController.Route(v1)
	foodTrxController.Route(v1)

	// Start App
	err := app.Listen(os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
