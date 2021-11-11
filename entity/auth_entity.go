package entity

type User struct {
	Id       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Role     string `gorm:"column:role;default:customer"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
