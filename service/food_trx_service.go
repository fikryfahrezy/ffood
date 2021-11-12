package service

import "github.com/fikryfahrezy/ffood/model"

type FoodTrxService interface {
	InsertFoodTrx(Request model.FoodTrxRequest, BuyerId string) (Response model.FoodTrxResponse, Error error)
}
