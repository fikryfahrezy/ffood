package repository

import (
	"golang-simple-boilerplate/entity"
	"golang-simple-boilerplate/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Mysql gorm.DB
}

func NewUserRepository(Mysql *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Mysql: *Mysql,
	}
}

func (Repository UserRepositoryImpl) Profile(Request model.ProfileRequest) (Response model.ProfileResponse, Error error) {
	var user entity.User
	if Error = Repository.Mysql.Where("email = ?", Request.Email).Find(&user).Error; Error != nil {
		return Response, Error
	}
	Response.Id = user.Id
	Response.Email = user.Email
	return
}

func (Repository UserRepositoryImpl) UpdateProfile(Request model.UpdateProfileRequest) (Response model.UpdateProfileResponse, Error error) {
	var user entity.User
	Error = Repository.Mysql.Model(&user).Where("email = ?", Request.Email).Updates(entity.User{
		Name:     Request.Name,
		Password: Request.Password,
	}).Error
	Response.Email = Request.Email
	Response.Name = Request.Name
	return
}
