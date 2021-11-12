package service

import (
	"github.com/fikryfahrezy/ffood/model"
	"github.com/fikryfahrezy/ffood/repository"
	"github.com/fikryfahrezy/ffood/validation"
)

type FoodServiceImpl struct {
	FoodRepository repository.FoodRepository
}

func NewFoodService(FoodRepository *repository.FoodRepository) FoodService {
	return &FoodServiceImpl{
		FoodRepository: *FoodRepository,
	}
}

func (Service FoodServiceImpl) InsertFood(Request model.InsertFoodRequest, SellerId int64) (Response model.InsertFoodResponse, Error error) {
	if Error = validation.InsertFoodValidation(Request); Error != nil {
		return
	}
	Response, Error = Service.FoodRepository.InsertFood(Request, SellerId)
	return
}

func (Service FoodServiceImpl) GetAllFood() (Response []model.InsertFoodResponse, Error error) {
	Response = Service.FoodRepository.GetAllFood()
	return
}

func (Service FoodServiceImpl) GetFood(Id string) (Response model.InsertFoodResponse, Error error) {
	Response, Error = Service.FoodRepository.GetFood(Id)
	return
}

func (Service FoodServiceImpl) DeleteFood(Id string, SellerId int64) (Response model.InsertFoodResponse, Error error) {
	Response, Error = Service.FoodRepository.DeleteFood(Id, SellerId)
	return
}

func (Service FoodServiceImpl) UpdateFood(Request model.UpdateFoodRequest, Id string, SellerId int64) (Response model.UpdateFoodResponse, Error error) {
	Response, Error = Service.FoodRepository.UpdateFood(Request, Id, SellerId)
	return
}
