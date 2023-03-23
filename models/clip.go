package models

type Clip struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	URL    string `json:"url"`
	UserID uint
	User   User
    Description string 
}
