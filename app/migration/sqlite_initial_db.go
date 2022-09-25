package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	model "github.com/joshuapohan/microapp/model"
)

func main() {
	db, err := gorm.Open(sqlite.Open("microapp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.LoginHistory{})
}
