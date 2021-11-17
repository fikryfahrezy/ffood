package repository

import "github.com/fikryfahrezy/ffood/model"

type FoodTrxRepository interface {
	InsertFoodTrx(Request model.InsertFoodTrxRequest, BuyerId int64) (Response model.FoodTrxResponse, Error error)
}
