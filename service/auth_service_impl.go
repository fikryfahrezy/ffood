package service

import (
	"errors"
	"time"

	"github.com/fikryfahrezy/ffood/helper"
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/validation"

	"github.com/dgrijalva/jwt-go"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
}

func NewAuthService(AuthRepository *repository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		AuthRepository: *AuthRepository,
	}
}

func (Service AuthServiceImpl) Register(Request model.RegisterRequest) (Response model.RegisterResponse, Error error) {
	if Error = validation.RegisterValidation(Request); Error != nil {
		return
	}

	hash, ok := helper.GenerateHash(Request.Password)
	if !ok {
		return
	}

	Request.Password = hash
	Response, Error = Service.AuthRepository.Register(Request)
	if Error != nil {
		return Response, errors.New("EMAIL_REGISTERED")
	}

	accessToken := helper.SignJWT(jwt.MapClaims{
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"id":    Response.Id,
		"email": Response.Email,
		"role":  Response.Role,
	})
	Response.AccessToken = accessToken
	return
}

func (Service AuthServiceImpl) Login(Request model.AuthRequest) (Response model.AuthResponse, Verified bool, Error error) {
	if Error = validation.LoginValidation(Request); Error != nil {
		return Response, Verified, Error
	}
	Response, userExists, Error := Service.AuthRepository.Login(Request)
	if userExists {
		if helper.CompareHash(Response.Password, Request.Password) {
			accessToken := helper.SignJWT(jwt.MapClaims{
				"exp":   time.Now().Add(24 * time.Hour).Unix(),
				"id":    Response.Id,
				"email": Response.Email,
				"role":  Response.Role,
			})
			Response.AccessToken = accessToken
			return Response, true, Error
		}
	}
	return Response, Verified, Error
}
