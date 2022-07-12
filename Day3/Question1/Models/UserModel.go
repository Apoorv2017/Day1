package Models

import "github.com/jinzhu/gorm"
type User struct {
	gorm.Model
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string{
	return "user"
}