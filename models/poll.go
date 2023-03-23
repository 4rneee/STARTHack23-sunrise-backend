package models

type Poll struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Question string `json:"question"`
}

type PollAnswer struct {
	ID     uint `json:"id" gorm:"primary_key"`
    Votes  uint `json:"votes"`
	PollID uint `json:"pollID"`
	Poll   Poll
}
