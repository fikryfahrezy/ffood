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

func (Service FoodServiceImpl) Insert(Request model.InsertFoodRequest, SellerId int64) (Response model.InsertFoodResponse, Error error) {
	if Error = validation.InsertFoodValidation(Request); Error != nil {
		return
	}
	Response, Error = Service.FoodRepository.Insert(Request, SellerId)
	return
}

func (Service FoodServiceImpl) GetAll() (Response []model.InsertFoodResponse, Error error) {
	Response = Service.FoodRepository.GetAll()
	return
}

func (Service FoodServiceImpl) Get(Id string) (Response model.InsertFoodResponse, Error error) {

	Response, Error = Service.FoodRepository.Get(Id)
	return
}

func (Service FoodServiceImpl) Delete(Id string, SellerId int64) (Response model.InsertFoodResponse, Error error) {
	Response, Error = Service.FoodRepository.Delete(Id, SellerId)
	return
}

func (Service FoodServiceImpl) Update(Request model.UpdateFoodRequest, Id string, SellerId int64) (Response model.UpdateFoodResponse, Error error) {
	Response, Error = Service.FoodRepository.Update(Request, Id, SellerId)
	return
}
