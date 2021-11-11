package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"golang-simple-boilerplate/model"
)

type DecodedStructure struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}

func DecodeToken(encodedToken string) (decodedResult DecodedStructure, errData error) {
	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return decodedResult, err
	}
	if !token.Valid {
		return decodedResult, errors.New("invalid token")
	}

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return decodedResult, err
	}

	var obj DecodedStructure
	if err := json.Unmarshal(jsonbody, &obj); err != nil {
		return decodedResult, err
	}

	return obj, nil
}

func CheckToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")

		var tokenString string
		if len(tokenSlice) == 2 {
			tokenString = tokenSlice[1]
		}

		// validate token
		_, err := ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err)
			response := model.Response{
				Code:   401,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// extract data from token
		decodedRes, err := DecodeToken(tokenString)
		if err != nil {
			fmt.Println(err)
			response := model.Response{
				Code:   401,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// set to global var
		c.Locals("id", decodedRes.Id)
		c.Locals("email", decodedRes.Email)
		return c.Next()
	}
}

func CheckRole(roles string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		getrole := c.Locals("currentRole").(string)
		roleArr := strings.Split(roles, ",")
		isOk := false
		for _, role := range roleArr {
			if role == getrole {
				isOk = true
			}
		}

		if !isOk {
			return c.Status(403).JSON(model.Response{
				Code:   403,
				Status: "Forbidden",
			})
		}

		return c.Next()
	}
}
