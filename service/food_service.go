package service

import "github.com/fikryfahrezy/ffood/model"

type FoodService interface {
	Insert(Request model.InsertFoodRequest, SellerId int64) (Response model.InsertFoodResponse, Error error)
	GetAll() (Response []model.InsertFoodResponse, Error error)
	Get(Id string) (Response model.InsertFoodResponse, Error error)
	Delete(Id string, SellerId int64) (Response model.InsertFoodResponse, Error error)
	Update(Request model.UpdateFoodRequest, Id string, SellerId int64) (Response model.UpdateFoodResponse, Error error)
}
