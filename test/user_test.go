package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/ffood/model"

	"github.com/fikryfahrezy/ffood/config"
	"github.com/fikryfahrezy/ffood/controller"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func userInit(gormDb *gorm.DB) (app *fiber.App, userService service.UserService, userRepository repository.UserRepository, authService service.AuthService) {
	authRepository := repository.NewAuthRepository(gormDb)
	authService = service.NewAuthService(&authRepository)

	userRepository = repository.NewUserRepository(gormDb)
	userService = service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	app = fiber.New(config.NewFiberConfig())
	userController.Route(app)
	return
}

func TestGetProfile(t *testing.T) {
	gormDb := dbInit()
	app, _, _, authService := userInit(gormDb)
	clearDb(gormDb)

	testCases := []struct {
		testName           string
		init               func(req *http.Request)
		method             string
		url                string
		expectedStatusCode int
	}{
		{
			testName: "Get Profile Success",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email8@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			method:             "GET",
			url:                "/user/profile",
			expectedStatusCode: http.StatusOK,
		},
		{
			testName:           "Get Profile Failed, No Token Provided",
			init:               func(req *http.Request) {},
			method:             "GET",
			url:                "/user/profile",
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			testName: "Get Profile Failed, Token Is Not Bearer Token",
			init: func(req *http.Request) {
				req.Header.Add("Authorization", fmt.Sprintf("Blablabla"))
			},
			method:             "POST",
			url:                "/user/profile",
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName: "Get Profile Failed, Bearer Token Is Not Valid",
			init: func(req *http.Request) {
				req.Header.Add("Authorization", fmt.Sprintf("Bearer Blablabla"))
			},
			method:             "POST",
			url:                "/user/profile",
			expectedStatusCode: http.StatusInternalServerError,
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

func TestUpdateProfile(t *testing.T) {
	gormDb := dbInit()
	app, userService, _, authService := userInit(gormDb)
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
			testName: "Update Profile Success, Update Name Success",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email9@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			after: func() bool {
				profile := model.ProfileRequest{
					Email: "email9@email.com",
				}
				response, _ := userService.Profile(profile)
				return response.Name == "New Name"
			},
			method:             "PATCH",
			url:                "/user/profile",
			body:               `{"name": "New Name"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Update Profile Success, Update Password Success",
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
			after: func() bool {
				auth := model.AuthRequest{Email: "email10@email.com", Password: "new_password"}
				_, isValid, _ := authService.Login(auth)

				return isValid
			},
			method:             "PATCH",
			url:                "/user/profile",
			body:               `{"password": "new_password"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Update Profile Success, Update Name & Password Success",
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
			after: func() bool {
				auth := model.AuthRequest{Email: "email11@email.com", Password: "new_password"}
				response, isValid, _ := authService.Login(auth)

				return isValid && response.Name == "New Name"
			},
			method:             "PATCH",
			url:                "/user/profile",
			body:               `{"name": "New Name", "password": "new_password"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Update Profile Success, Empty Request",
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
			after: func() bool {
				auth := model.AuthRequest{Email: "email12@email.com", Password: "password"}
				response, isValid, _ := authService.Login(auth)

				fmt.Println(response.Name, isValid)
				return isValid && response.Name == "Name"
			},
			method:             "PATCH",
			url:                "/user/profile",
			body:               `{}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Update Profile Fail, Request Not JSON",
			init: func(req *http.Request) {
				user := model.RegisterRequest{
					Email:    "email13@email.com",
					Name:     "Name",
					Password: "password",
				}
				response, _ := authService.Register(user)

				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AccessToken))
			},
			after: func() bool {
				return true
			},
			method:             "PATCH",
			url:                "/user/profile",
			body:               ``,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName: "Update Profile Fail, No Token Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			after: func() bool {
				return true
			},
			method:             "PATCH",
			url:                "/user/profile",
			body:               `{}`,
			expectedStatusCode: http.StatusUnauthorized,
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
