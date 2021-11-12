package repository

import (
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

func (Repository FoodTrxRepositoryImpl) InsertFoodTrx(Request model.FoodTrxRequest, BuyerId int64) (Response model.FoodTrxResponse, Error error) {
	return
}
