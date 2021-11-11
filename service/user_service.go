package service

import "github.com/fikryfahrezy/ffood/model"

type UserService interface {
	Profile(Request model.ProfileRequest) (Response model.ProfileResponse, Error error)
	UpdateProfile(Request model.UpdateProfileRequest) (Response model.UpdateProfileResponse, Error error)
}
