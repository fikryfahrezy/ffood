package entity

type Food struct {
	Id        int64  `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	IsDeleted bool   `gorm:"column:is_deleted;default:0"`
	SellerId  int64  `gorm:"column:seller_id"`
	Seller    User   `gorm:"foreignKey:SellerId;references:Id"`
}

func (Food) TableName() string {
	return "foods"
}
