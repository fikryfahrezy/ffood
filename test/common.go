package test

import (
	"database/sql"
	"os"

	"github.com/fikryfahrezy/ffood/exception"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbInit() (gormDb *gorm.DB) {
	var err error
	err = godotenv.Load("../.env.test")
	exception.PanicIfNeeded(err)

	sqlDB, err := sql.Open("mysql", os.Getenv("MYSQL_HOST"))
	exception.PanicIfNeeded(err)

	gormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	exception.PanicIfNeeded(err)

	return
}

func clearDb(mysql *gorm.DB) {
	mysql.Exec("DELETE FROM food_transactions")
	mysql.Exec("DELETE FROM foods")
	mysql.Exec("DELETE FROM users")
}
