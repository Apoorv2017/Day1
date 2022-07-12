package main

import (
	"question2/Models"
	
	"question2/Controllers"
	
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
		grp.GET("get-all-students", Controllers.GetUsers)
		grp.POST("insert-student", Controllers.CreateUser)
		grp.GET("get-student/:id", Controllers.GetUserByID)
		grp.PUT("update-student/:id", Controllers.UpdateUser)
		grp.DELETE("delete-student/:id", Controllers.DeleteUser)
		
	}
	
	router.Run()
	
}