package repository

import (
	"golang-simple-boilerplate/entity"
	"golang-simple-boilerplate/model"

	"gorm.io/gorm"
)

type FoodRepositoryImpl struct {
	Mysql gorm.DB
}

func NewFoodRepository(Mysql *gorm.DB) FoodRepository {
	return &FoodRepositoryImpl{
		Mysql: *Mysql,
	}
}

func (Repository FoodRepositoryImpl) Insert(Request model.FoodRequest, SellerId int64) (Response model.FoodResponse, Error error) {
	food := entity.Food{
		Name:     Request.Name,
		SellerId: SellerId,
	}
	Error = Repository.Mysql.Create(&food).Error
	if Error != nil {
		return
	}

	Response.Id = food.Id
	Response.Name = food.Name
	Response.SellerId = food.SellerId
	Response.Seller = food.Seller
	return
}

func (Repository FoodRepositoryImpl) GetAll() (Response []model.FoodResponse) {
	var foods []entity.Food
	Repository.Mysql.Where("is_deleted = ?", nil).Find(&foods)

	Response = make([]model.FoodResponse, len(foods))
	for i, food := range foods {
		Response[i] = model.FoodResponse{
			Id:       food.Id,
			Name:     food.Name,
			SellerId: food.SellerId,
			Seller:   food.Seller,
		}
	}
	return
}

func (Repository FoodRepositoryImpl) Get(Id string) (Response model.FoodResponse, Error error) {
	var food entity.Food
	Error = Repository.Mysql.First(&food, "id = ? AND is_deleted = ?", Id, nil).Error
	if Error != nil {
		return
	}
	Response.Id = food.Id
	Response.Name = food.Name
	Response.SellerId = food.SellerId
	Response.Seller = food.Seller
	return
}

func (Repository FoodRepositoryImpl) Delete(Id string, SellerId int64) (Response model.FoodResponse, Error error) {
	var food entity.Food
	Error = Repository.Mysql.First(&food, "id = ? AND is_deleted = ? AND seller_id = ?", Id, nil, SellerId).Error
	if Error != nil {
		return
	}

	Error = Repository.Mysql.Save(&food).Error
	if Error != nil {
		return
	}
	Response.Id = food.Id
	Response.Name = food.Name
	Response.SellerId = food.SellerId
	Response.Seller = food.Seller
	return
}
