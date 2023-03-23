package models

type Comment struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Content  string `json:"content"`
	UserID   uint   `json:"userid"`
	User     User
	StreamID uint `json:"streamid"`
	Stream   Stream
}
