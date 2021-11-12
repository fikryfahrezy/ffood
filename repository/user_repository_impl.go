package repository

import (
	"github.com/fikryfahrezy/ffood/entity"
	"github.com/fikryfahrezy/ffood/model"

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
	Response.Name = user.Name
	Response.Role = user.Role
	return
}

func (Repository UserRepositoryImpl) UpdateProfile(Request model.UpdateProfileRequest) (Response model.UpdateProfileResponse, Error error) {
	var user entity.User
	Error = Repository.Mysql.Where("email = ?", Request.Email).First(&user).Error
	if Error != nil {
		return
	}

	if Request.Name != "" {
		user.Name = Request.Name
	}
	if Request.Password != "" {
		user.Password = Request.Password
	}

	Error = Repository.Mysql.Save(&user).Error
	if Error != nil {
		return
	}

	Response.Id = user.Id
	Response.Email = user.Email
	Response.Name = user.Name
	Response.Role = user.Role
	return
}
