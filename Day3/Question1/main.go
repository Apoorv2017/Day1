package main

import (
	"gorm+gim/Models"
	
	"gorm+gim/Controllers"
	
	// "github.com/jinzhu/gorm"
//   _ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error
var router *gin.Engine


func main(){
	
  	db, err:= gorm.Open(sqlite.Open("Database.db"), &gorm.Config{})
	if err!=nil{
		panic("gailed to connect database")
	}
	db.AutoMigrate(&Models.User{})
	
	router := gin.Default()
	grp:=router.Group("/user-api")
	{
		grp.GET("user", Controllers.GetUsers)
		grp.POST("user", Controllers.CreateUser)
		grp.GET("user/:id", Controllers.GetUserByID)
		grp.PUT("user/:id", Controllers.UpdateUser)
		grp.DELETE("user/:id", Controllers.DeleteUser)
		
	}
	
	router.Run()
	
}