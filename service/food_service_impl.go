package service

import (
	"golang-simple-boilerplate/model"
	"golang-simple-boilerplate/repository"
)

type FoodServiceImpl struct {
	FoodRepository repository.FoodRepository
}

func NewFoodService(FoodRepository *repository.FoodRepository) FoodService {
	return &FoodServiceImpl{
		FoodRepository: *FoodRepository,
	}
}

func (Service FoodServiceImpl) Insert(Request model.FoodRequest, SellerId string) (Response model.FoodResponse, Error error) {
	Response, Error = Service.FoodRepository.Insert(Request, 0)
	return
}

func (Service FoodServiceImpl) GetAll() (Response []model.FoodResponse, Error error) {
	Response = Service.FoodRepository.GetAll()
	return
}

func (Service FoodServiceImpl) Get(Id string) (Response model.FoodResponse, Error error) {
	Response, Error = Service.FoodRepository.Get(Id)
	return
}

func (Service FoodServiceImpl) Delete(Id string, SellerId string) (Response model.FoodResponse, Error error) {
	Response, Error = Service.FoodRepository.Delete(Id, 0)
	return
}
