package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/ffood/entity"

	"github.com/fikryfahrezy/ffood/model"

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
	app, _, _, authService := foodTrxInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		body               string
		expectedStatusCode int
	}{
		{
			testName: "Insert Food Trx Success",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email9@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				food := entity.Food{
					Id:       1,
					Name:     "Name",
					SellerId: response.Id,
				}
				gormDb.Create(&food)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foodtransactions",
			body:               `{"food_id": 1}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Insert Food Trx Fail, Food Id Not Found",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email10@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foodtransactions",
			body:               `{"food_id": 2}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName: "Insert Food Trx Fail, No Token Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/foodtransactions",
			body:               `{"food_id": 2}`,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			testName: "Insert Food Trx Fail, Request Empty JSON",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email11@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foodtransactions",
			body:               `{}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Insert Food Trx Fail, Request Empty Body",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email12@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foodtransactions",
			body:               ``,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

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
