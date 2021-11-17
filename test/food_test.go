package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/ffood/entity"

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
					Email:    "email1@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email1@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email1@email.com",
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
		{
			testName: "Insert Food Fail, Empty Body JSON",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email2@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email2@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email2@email.com",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foods",
			body:               `{}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Insert Food Fail, Empty Body",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email3@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email3@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email3@email.com",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foods",
			body:               ``,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName: "Insert Food Fail, No Token Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/foods",
			body:               `{}`,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			testName: "Insert Food Fail, User Role Not Allowed",
			init: func(req *http.Request) {
				response, _ := authService.Register(model.RegisterRequest{
					Email:    "email4@email.com",
					Name:     "Name",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "POST",
			url:                "/foods",
			body:               `{}`,
			expectedStatusCode: http.StatusForbidden,
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
	app, _, _, authService := foodInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		expectedStatusCode int
	}{
		{
			testName: "Get Food Success",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email1@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email1@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email1@email.com",
					Password: "password",
				})

				food := entity.Food{
					Id:       1,
					Name:     "Food Name",
					SellerId: response.Id,
				}
				gormDb.Create(&food)

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "GET",
			url:                "/foods/1",
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Get Food Fail, Food Not Found",
			init: func(req *http.Request) {
				response, _ := authService.Register(model.RegisterRequest{
					Email:    "email2@email.com",
					Name:     "Name",
					Password: "password",
				})

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "GET",
			url:                "/foods/9999",
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
	app, _, _, authService := foodInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		expectedStatusCode int
	}{
		{
			testName: "Delete Food Success",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email1@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email1@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email1@email.com",
					Password: "password",
				})

				food := entity.Food{
					Id:       1,
					Name:     "Food Name",
					SellerId: response.Id,
				}
				gormDb.Create(&food)

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "DELETE",
			url:                "/foods/1",
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Delete Food Fail, Food Not Found",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email2@email.com",
					Name:     "Name",
					Password: "password",
				})

				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email2@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email2@email.com",
					Password: "password",
				})

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "DELETE",
			url:                "/foods/9999",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			testName:           "Delete Food Fail, No Token Provided",
			init:               func(req *http.Request) {},
			method:             "DELETE",
			url:                "/foods/9999",
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			testName: "Insert Food Fail, User Role Not Allowed",
			init: func(req *http.Request) {
				response, _ := authService.Register(model.RegisterRequest{
					Email:    "email3@email.com",
					Name:     "Name",
					Password: "password",
				})

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "DELETE",
			url:                "/foods/999",
			expectedStatusCode: http.StatusForbidden,
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

func TestUpdateFood(t *testing.T) {
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
			testName: "Update Food Success",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email1@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email1@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email1@email.com",
					Password: "password",
				})

				food := entity.Food{
					Id:       1,
					Name:     "Food Name",
					SellerId: response.Id,
				}
				gormDb.Create(&food)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "PATCH",
			url:                "/foods/1",
			body:               `{"name": "New Food Name"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Update Food Fail, Food Not Found",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email2@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email2@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email2@email.com",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "PATCH",
			url:                "/foods/999",
			body:               `{"name": "New Food Name"}`,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			testName: "Update Food Success, Empty Body JSON",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email3@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email3@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email3@email.com",
					Password: "password",
				})

				food := entity.Food{
					Id:       9,
					Name:     "Food Name",
					SellerId: response.Id,
				}
				gormDb.Create(&food)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "PATCH",
			url:                "/foods/9",
			body:               `{}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Update Food Fail, Empty Body",
			init: func(req *http.Request) {
				authService.Register(model.RegisterRequest{
					Email:    "email4@email.com",
					Name:     "Name",
					Password: "password",
				})
				gormDb.Exec("UPDATE users SET role = 'seller' WHERE email = 'email4@email.com'")
				response, _, _ := authService.Login(model.AuthRequest{
					Email:    "email4@email.com",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "PATCH",
			url:                "/foods/999",
			body:               ``,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName: "Update Food Fail, No Token Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "PATCH",
			url:                "/foods/999",
			body:               `{}`,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			testName: "Patch Food Fail, User Role Not Allowed",
			init: func(req *http.Request) {
				response, _ := authService.Register(model.RegisterRequest{
					Email:    "email5@email.com",
					Name:     "Name",
					Password: "password",
				})

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "PATCH",
			url:                "/foods/999",
			body:               `{}`,
			expectedStatusCode: http.StatusForbidden,
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
