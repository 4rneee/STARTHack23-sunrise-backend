package models

type Tweet struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserName string `json:"username"`
	Content  string `json:"content"`
}
