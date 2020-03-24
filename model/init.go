package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/crud?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	DB = db

	db.AutoMigrate(&User{}, &Article{})
}
