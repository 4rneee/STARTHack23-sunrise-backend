package models

import "fmt"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `gorm:"type:BINARY(60)"`
	Points   uint
	Friends  []User `gorm:"many2many:friends;association_jointable_foreignkey:friend_id"`
	StreamID uint   
}

func (u *User) String() string {
	return fmt.Sprintf("{%v %v %v, %v, %v}", u.ID, u.Name, u.Email, string(u.Password), u.Points)
}
