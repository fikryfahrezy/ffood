package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/ffood/model"

	"gorm.io/gorm"

	"github.com/fikryfahrezy/ffood/config"
	"github.com/fikryfahrezy/ffood/controller"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/service"
	"github.com/gofiber/fiber/v2"
)

func authInit(gormDb *gorm.DB) (app *fiber.App, authService service.AuthService, authRepository repository.AuthRepository) {
	authRepository = repository.NewAuthRepository(gormDb)
	authService = service.NewAuthService(&authRepository)
	authController := controller.NewAuthController(&authService)

	app = fiber.New(config.NewFiberConfig())
	authController.Route(app)
	return
}

func TestRegister(t *testing.T) {
	gormDb := dbInit()
	app, _, authRepository := authInit(gormDb)
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
			testName: "Register Success",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{"email": "email1@email.com", "password": "password", "name": "Name"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Register Fail, Email Registered",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
				user := model.RegisterRequest{
					Email:    "email2@email.com",
					Name:     "Name",
					Password: "password",
				}
				authRepository.Register(user)
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{"email": "email2@email.com", "password": "password", "name": "Name"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Register Fail, Email Not Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{"password": "wrong_password", "name": "name"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Register Fail, Email Format Is Not Valid",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{"email": "not_email_format", "name": "name"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Register Fail, Password Not Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{"email": "email3@email.com", "name": "name"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Register Fail, Name Not Provided",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{"email": "email4@email.com", "password": "password"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Register Fail, Body Is Empty JSON",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
			body:               `{}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "Register Fail, Body Is Empty",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/register",
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

func TestLogin(t *testing.T) {
	gormDb := dbInit()
	app, authService, _ := authInit(gormDb)
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
			testName: "Login Success",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
				user := model.RegisterRequest{
					Email:    "email5@email.com",
					Name:     "Name",
					Password: "password",
				}
				authService.Register(user)
			},
			method:             "POST",
			url:                "/auth/login",
			body:               `{"email": "email5@email.com", "password": "password"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			testName: "Login Fail, User Not Registered",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
			},
			method:             "POST",
			url:                "/auth/login",
			body:               `{"email": "email6@email.com", "password": "password"}`,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			testName: "Login Fail, Wrong Password",
			init: func(req *http.Request) {
				req.Header.Add("Content-Type", "application/json")
				user := model.RegisterRequest{
					Email:    "email7@email.com",
					Name:     "Name",
					Password: "password",
				}
				authService.Register(user)
			},
			method:             "POST",
			url:                "/auth/login",
			body:               `{"email": "email7@email.com", "password": "wrong_password"}`,
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
		})
	}
}
