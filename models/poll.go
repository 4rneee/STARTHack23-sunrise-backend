package models

import "github.com/jinzhu/gorm"

type Poll struct {
	gorm.Model
	Question string `json:"question"`
	StreamID uint   `json:"streamid"`
	Stream   Stream
}

type PollAnswer struct {
	ID     uint `json:"id" gorm:"primary_key"`
	Votes  uint `json:"votes"`
	PollID uint `json:"pollID"`
	Poll   Poll
	Answer string `json:"answer"`
}
