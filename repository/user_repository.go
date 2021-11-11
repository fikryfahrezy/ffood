package repository

import "golang-simple-boilerplate/model"

type UserRepository interface {
	Profile(Request model.ProfileRequest) (Response model.ProfileResponse, Error error)
	UpdateProfile(Request model.UpdateProfileRequest) (Response model.UpdateProfileResponse, Error error)
}
