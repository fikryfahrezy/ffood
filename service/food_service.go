package service

import "golang-simple-boilerplate/model"

type FoodService interface {
	Insert(Request model.FoodRequest, SellerId string) (Response model.FoodResponse, Error error)
	GetAll() (Response []model.FoodResponse, Error error)
	Get(Id string) (Response model.FoodResponse, Error error)
	Delete(Id string, SellerId string) (Response model.FoodResponse, Error error)
}
