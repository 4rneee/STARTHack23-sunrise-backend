package models

import "github.com/jinzhu/gorm"

type Poll struct {
    gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Question string `json:"question"`
}

type PollAnswer struct {
	ID     uint `json:"id" gorm:"primary_key"`
    Votes  uint `json:"votes"`
	PollID uint `json:"pollID"`
	Poll   Poll
}
