package helper

import (
	"github.com/fikryfahrezy/ffood/exception"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func SignJWT(claims jwt.MapClaims) string {
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	exception.PanicIfNeeded(err)
	return accessToken
}
