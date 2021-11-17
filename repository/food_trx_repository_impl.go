package repository

import (
	"github.com/fikryfahrezy/ffood/entity"
	"github.com/fikryfahrezy/ffood/model"
	"gorm.io/gorm"
)

type FoodTrxRepositoryImpl struct {
	Mysql gorm.DB
}

func NewFoodTrxRepository(Mysql *gorm.DB) FoodTrxRepository {
	return &FoodTrxRepositoryImpl{
		Mysql: *Mysql,
	}
}

func (Repository FoodTrxRepositoryImpl) InsertFoodTrx(Request model.InsertFoodTrxRequest, BuyerId int64) (Response model.FoodTrxResponse, Error error) {
	foodTrx := entity.FoodTrx{
		FoodId:  Request.FoodId,
		BuyerId: BuyerId,
	}
	if Error = Repository.Mysql.Preload("Food.Seller").Create(&foodTrx).Error; Error != nil {
		return Response, Error
	}

	Response.Id = foodTrx.Id
	Response.Name = foodTrx.Food.Name
	Response.SellerId = foodTrx.Food.SellerId
	Response.Seller = foodTrx.Food.Seller.Name

	return Response, Error
}
