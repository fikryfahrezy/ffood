package model

type FoodTrxResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	SellerId int    `json:"seller_id"`
	Seller   string `json:"seller"`
}
