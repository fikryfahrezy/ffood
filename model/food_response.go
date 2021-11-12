package model

type InsertFoodResponse struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name"`
	SellerId int64       `json:"seller_id"`
	Seller   interface{} `json:"seller"`
}

type UpdateFoodResponse struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name"`
	SellerId int64       `json:"seller_id"`
	Seller   interface{} `json:"seller"`
}
