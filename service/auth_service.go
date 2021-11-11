package service

import "golang-simple-boilerplate/model"

type AuthService interface {
	Register(Request model.RegisterRequest) (Response model.RegisterResponse, Error error)
	Login(Request model.AuthRequest) (Response model.AuthResponse, UserExists bool, Error error)
}
