package service

import (
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/repository"
)

type FoodTrxServiceImpl struct {
	FoodTrxRepository repository.FoodTrxRepository
}

func NewFoodTrxService(FoodTrxRepository *repository.FoodTrxRepository) FoodTrxService {
	return &FoodTrxServiceImpl{
		FoodTrxRepository: *FoodTrxRepository,
	}
}

func (Service FoodTrxServiceImpl) InsertFoodTrx(Request model.FoodTrxRequest, BuyerId string) (Response model.FoodTrxResponse, Error error) {
	return
}
