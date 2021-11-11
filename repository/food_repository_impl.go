package repository

import (
	"github.com/fikryfahrezy/ffood/entity"
	"github.com/fikryfahrezy/ffood/model"

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

func (Repository FoodRepositoryImpl) Insert(Request model.InsertFoodRequest, SellerId int64) (Response model.InsertFoodResponse, Error error) {
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
	//Response.Seller = food.Seller
	return
}

func (Repository FoodRepositoryImpl) GetAll() (Response []model.InsertFoodResponse) {
	var foods []entity.Food
	Repository.Mysql.Where("is_deleted = ?", 0).Find(&foods)

	Response = make([]model.InsertFoodResponse, len(foods))
	for i, food := range foods {
		Response[i] = model.InsertFoodResponse{
			Id:       food.Id,
			Name:     food.Name,
			SellerId: food.SellerId,
			//Seller:   food.Seller,
		}
	}
	return
}

func (Repository FoodRepositoryImpl) Get(Id string) (Response model.InsertFoodResponse, Error error) {
	var food entity.Food
	Error = Repository.Mysql.First(&food, "id = ? AND is_deleted = ?", Id, 0).Error
	if Error != nil {
		return
	}
	Response.Id = food.Id
	Response.Name = food.Name
	Response.SellerId = food.SellerId
	//Response.Seller = food.Seller
	return
}

func (Repository FoodRepositoryImpl) Delete(Id string, SellerId int64) (Response model.InsertFoodResponse, Error error) {
	var food entity.Food
	Error = Repository.Mysql.First(&food, "id = ? AND seller_id = ? AND is_deleted = ?", Id, SellerId, 0).Error
	if Error != nil {
		return
	}

	food.IsDeleted = true
	Error = Repository.Mysql.Save(&food).Error
	if Error != nil {
		return
	}
	Response.Id = food.Id
	Response.Name = food.Name
	Response.SellerId = food.SellerId
	//Response.Seller = food.Seller
	return
}

func (Repository FoodRepositoryImpl) Update(Request model.UpdateFoodRequest, Id string, SellerId int64) (Response model.UpdateFoodResponse, Error error) {
	var food entity.Food
	Error = Repository.Mysql.First(&food, "id = ? AND seller_id = ? AND is_deleted = ?", Id, SellerId, 0).Error
	if Error != nil {
		return
	}

	food.Name = Request.Name
	Error = Repository.Mysql.Save(&food).Error
	if Error != nil {
		return
	}
	Response.Id = food.Id
	Response.Name = food.Name
	Response.SellerId = food.SellerId
	//Response.Seller = food.Seller
	return
}
