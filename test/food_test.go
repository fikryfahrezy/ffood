package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/ffood/config"
	"github.com/fikryfahrezy/ffood/controller"
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func foodInit(gormDb *gorm.DB) (app *fiber.App, foodService service.FoodService, foodRepository repository.FoodRepository, authService service.AuthService) {
	authRepository := repository.NewAuthRepository(gormDb)
	authService = service.NewAuthService(&authRepository)

	foodRepository = repository.NewFoodRepository(gormDb)
	foodService = service.NewFoodService(&foodRepository)
	foodController := controller.NewFoodController(&foodService)

	app = fiber.New(config.NewFiberConfig())
	foodController.Route(app)
	return
}

func TestInsertFood(t *testing.T) {
	gormDb := dbInit()
	app, _, _, authService := foodInit(gormDb)
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
			testName: "Insert Food Success",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email14@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email14@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email14@email.com",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foods",
			body:               `{"name": "Food Name"}`,
			expectedStatusCode: http.StatusOK,
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

func TestGetFoods(t *testing.T) {
	gormDb := dbInit()
	app, _, _, _ := foodInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		expectedStatusCode int
	}{
		{
			testName:           "Get All Food Success",
			init:               func(req *http.Request) {},
			method:             "GET",
			url:                "/foods",
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			req := httptest.NewRequest(testCase.method, testCase.url, nil)
			testCase.init(req)

			resp, _ := app.Test(req)
			if resp.StatusCode != testCase.expectedStatusCode {
				body, _ := ioutil.ReadAll(resp.Body)
				t.Fatal(string(body))
			}
		})
	}
}

func TestGetFood(t *testing.T) {
	gormDb := dbInit()
	app, _, _, _ := foodInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		expectedStatusCode int
	}{
		{
			testName:           "Food Fail, Food Not Found",
			init:               func(req *http.Request) {},
			method:             "GET",
			url:                "/food/1",
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			req := httptest.NewRequest(testCase.method, testCase.url, nil)
			testCase.init(req)

			resp, _ := app.Test(req)
			if resp.StatusCode != testCase.expectedStatusCode {
				body, _ := ioutil.ReadAll(resp.Body)
				t.Fatal(string(body))
			}
		})
	}
}

func TestDeleteFood(t *testing.T) {
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
	}{
		{
			testName:           "Delete Food Fail, Food Not Found",
			init:               func(req *http.Request) {},
			method:             "DELETE",
			url:                "/food/1",
			body:               `{}`,
			expectedStatusCode: http.StatusNotFound,
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

func TestUpdateFood(t *testing.T) {
	gormDb := dbInit()
	app, _, _, _ := foodInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		after              func() bool
		method             string
		url                string
		body               string
		expectedStatusCode int
	}{
		{
			testName: "Update Food Fail, Food Not Found",
			init:     func(req *http.Request) {},
			after: func() bool {
				return true
			},
			method:             "PATCH",
			url:                "/food/1",
			body:               `{}`,
			expectedStatusCode: http.StatusNotFound,
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

			if !testCase.after() {
				t.Fatal("After result is false")
			}
		})
	}
}
