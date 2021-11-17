package service

import "github.com/fikryfahrezy/ffood/model"

type FoodTrxService interface {
	InsertFoodTrx(Request model.InsertFoodTrxRequest, BuyerId int64) (Response model.FoodTrxResponse, Error error)
}
