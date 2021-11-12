package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/ffood/config"
	"github.com/fikryfahrezy/ffood/controller"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func foodTrxInit(gormDb *gorm.DB) (app *fiber.App, foodTrxService service.FoodTrxService, foodTrxRepository repository.FoodTrxRepository, authService service.AuthService) {
	authRepository := repository.NewAuthRepository(gormDb)
	authService = service.NewAuthService(&authRepository)

	foodTrxRepository = repository.NewFoodTrxRepository(gormDb)
	foodTrxService = service.NewFoodTrxService(&foodTrxRepository)
	foodTrxController := controller.NewFoodTrxController(&foodTrxService)

	app = fiber.New(config.NewFiberConfig())
	foodTrxController.Route(app)
	return
}

func TestInsertFoodTrx(t *testing.T) {
	gormDb := dbInit()
	app, _, _, _ := foodInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		body               string
		expectedStatusCode int
	}{}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			req := httptest.NewRequest(testCase.method, testCase.url, strings.NewReader(testCase.body))
			testCase.init(req)

			resp, _ := app.Test(req)
			if resp.StatusCode != testCase.expectedStatusCode {
				body, _ := ioutil.ReadAll(resp.Body)
				t.Fatal(string(body))
			}
		})
	}
}
