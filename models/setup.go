package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	database, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{}, &User{})
	if err != nil {
		return
	}

	DB  = database
}