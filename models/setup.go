package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "sunrise.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.LogMode(true)

	database.AutoMigrate(&User{}, &Poll{}, &PollAnswer{}, &Stream{}, &Comment{}, &Clip{})

	DB = database
}
