package service

import "github.com/fikryfahrezy/ffood/model"

type AuthService interface {
	Register(Request model.RegisterRequest) (Response model.RegisterResponse, Error error)
	Login(Request model.AuthRequest) (Response model.AuthResponse, UserExists bool, Error error)
}
