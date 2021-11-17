package model

type FoodTrxResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	SellerId int64  `json:"seller_id"`
	Seller   string `json:"seller"`
}
