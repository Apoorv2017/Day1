package Models
import (
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)


func GetAllUsers(user *[]User) (err error) {
	db, err := gorm.Open("sqlite3", "Database.db")
	if err!=nil{
		panic("error in getting all user")
	}
	
	if err = db.Find(user).Error; err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}


func CreateUser(user *User) (err error) {
	db, err := gorm.Open("sqlite3", "Database.db")
	if err!=nil{
		panic("error while creating user")
	}
	
	if err = db.Create(user).Error; err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}


func GetUserByID(user *User, id string) (err error) {
	db, err := gorm.Open("sqlite3", "Database.db")
	if err!=nil{
		panic("error while getting user by id")
	}
	
	if err = db.Where("id = ?", id).First(user).Error; err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}


func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	db, err := gorm.Open("sqlite3", "Database.db")
	if err!=nil{
		panic("error while updating user")
	}
	db.Save(user)
	db.Close()
	return nil
}


func DeleteUser(user *User, id string) (err error) {
	db, err := gorm.Open("sqlite3", "Database.db")
	if err!=nil{
		panic("error while deleting user")
	}
	db.Where("id = ?", id).Delete(user)
	db.Close()
	return nil
}