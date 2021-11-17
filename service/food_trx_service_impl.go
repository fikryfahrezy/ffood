package service

import (
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/validation"
)

type FoodTrxServiceImpl struct {
	FoodTrxRepository repository.FoodTrxRepository
}

func NewFoodTrxService(FoodTrxRepository *repository.FoodTrxRepository) FoodTrxService {
	return &FoodTrxServiceImpl{
		FoodTrxRepository: *FoodTrxRepository,
	}
}

func (Service FoodTrxServiceImpl) InsertFoodTrx(Request model.InsertFoodTrxRequest, BuyerId int64) (Response model.FoodTrxResponse, Error error) {
	if Error = validation.InsertFoodTrxValidation(Request); Error != nil {
		return
	}
	Response, Error = Service.FoodTrxRepository.InsertFoodTrx(Request, BuyerId)
	return
}
