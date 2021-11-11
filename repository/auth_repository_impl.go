package repository

import (
	"golang-simple-boilerplate/entity"
	"golang-simple-boilerplate/model"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Mysql gorm.DB
}

func NewAuthRepository(Mysql *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		Mysql: *Mysql,
	}
}

func (Repository AuthRepositoryImpl) Login(Request model.AuthRequest) (Response model.AuthResponse, UserExists bool, Error error) {
	var user entity.User
	Error = Repository.Mysql.Where("email = ?", Request.Email).Find(&user).Error
	if user.Email == "" {
		return Response, false, Error
	}
	Response.Id = user.Id
	Response.Name = user.Name
	Response.Email = user.Email
	Response.Password = user.Password
	Response.Role = user.Role
	return Response, true, Error
}

func (Repository AuthRepositoryImpl) Register(Request model.RegisterRequest) (Response model.RegisterResponse, Error error) {
	user := entity.User{
		Email:    Request.Email,
		Name:     Request.Name,
		Password: Request.Password,
	}
	Error = Repository.Mysql.Create(&user).Error
	if Error != nil {
		return
	}
	Response.Id = user.Id
	Response.Name = user.Name
	Response.Email = user.Email
	Response.Role = user.Role
	return
}
