package repository

import "github.com/fikryfahrezy/ffood/model"

type FoodRepository interface {
	InsertFood(Request model.InsertFoodRequest, SellerId int64) (Response model.InsertFoodResponse, Error error)
	GetAllFood() (Response []model.InsertFoodResponse)
	GetFood(Id string) (Response model.InsertFoodResponse, Error error)
	DeleteFood(Id string, SellerId int64) (Response model.InsertFoodResponse, Error error)
	UpdateFood(Request model.UpdateFoodRequest, Id string, SellerId int64) (Response model.UpdateFoodResponse, Error error)
}
