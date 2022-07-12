package Models

import "github.com/jinzhu/gorm"
type User struct {
	gorm.Model
	Id uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DOB string `json:"dob"`
	Address string `json:"address"`
	Subject string `json:"subject"`
	Marks string `json:"marks"`
}

func (b *User) TableName() string{
	return "user"
}