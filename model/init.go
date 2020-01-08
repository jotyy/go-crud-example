package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		fmt.Println(err)
	}

	DB = db

	db.AutoMigrate(&User{}, &Article{})
}
