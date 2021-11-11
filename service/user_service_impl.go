package service

import (
	"errors"
	"github.com/fikryfahrezy/ffood/helper"
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(UserRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: *UserRepository,
	}
}

func (Service UserServiceImpl) Profile(Request model.ProfileRequest) (Response model.ProfileResponse, Error error) {
	Response, Error = Service.UserRepository.Profile(Request)
	return Response, Error
}

func (Service UserServiceImpl) UpdateProfile(Request model.UpdateProfileRequest) (Response model.UpdateProfileResponse, Error error) {
	user, _ := Service.UserRepository.Profile(model.ProfileRequest{Email: Request.Email})
	if user.Id == 0 {
		return Response, errors.New("EMAIL_REGISTERED")
	}

	if Request.Password != "" {
		hash, ok := helper.GenerateHash(Request.Password)
		if !ok {
			return Response, errors.New("INTERNAL_SERVER_ERROR")
		}
		Request.Password = hash
	}

	Response, Error = Service.UserRepository.UpdateProfile(Request)
	return
}
