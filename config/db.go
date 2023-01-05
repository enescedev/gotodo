package config

import (
	"github.com/enescedev/gotodo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate($models.User{})
	DB=db
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

}