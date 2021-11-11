package model

type InsertFoodRequest struct {
	Name string `json:"name"`
}

type UpdateFoodRequest struct {
	Name string `json:"name"`
}
