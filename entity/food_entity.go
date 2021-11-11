package entity

import (
	"time"
)

type Food struct {
	Id        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	IsDeleted time.Time `gorm:"column:is_deleted;default:null"`
	SellerId  int64     `gorm:"column:seller_id"`
	Seller    User      `gorm:"foreignKey:SellerId;references:Id"`
}

func (Food) TableName() string {
	return "foods"
}
