package entity

type FoodTrx struct {
	Id      int64 `gorm:"column:id"`
	FoodId  int64 `gorm:"column:food_id"`
	BuyerId int64 `gorm:"column:buyer_id"`
	Food    Food  `gorm:"foreignKey:FoodId;references:Id"`
	Buyer   User  `gorm:"foreignKey:BuyerId;references:Id"`
}

func (FoodTrx) TableName() string {
	return "food_transactions"
}
