package repository

import "golang-simple-boilerplate/model"

type FoodRepository interface {
	Insert(Request model.FoodRequest, SellerId int64) (Response model.FoodResponse, Error error)
	GetAll() (Response []model.FoodResponse)
	Get(Id string) (Response model.FoodResponse, Error error)
	Delete(Id string, SellerId int64) (Response model.FoodResponse, Error error)
}
