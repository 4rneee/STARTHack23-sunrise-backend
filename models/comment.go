package models

import "github.com/jinzhu/gorm"

type Comment struct {
    gorm.Model
	Content  string `json:"content"`
	UserID   uint   `json:"userid"`
	User     User
	StreamID uint `json:"streamid"`
	Stream   Stream
}
