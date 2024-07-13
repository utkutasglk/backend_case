package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(){

	dsn := "hbstudent:hbstudent@tcp(127.0.0.1:3306)/backend_case?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB{
	return db
}