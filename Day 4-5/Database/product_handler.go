package Database

import (
	"project/Models"
	"project/Config"
	
	// "github.com/jinzhu/gorm"
//   _ "github.com/jinzhu/gorm/dialects/mysql"
 
  
  	
)



func GetProductById(prod *Models.Product, id string) (err error) {
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("error while getting user by id")
	// }
	
	if err = Config.DB.Where("product_id = ?", id).First(prod).Error; err != nil {
		// Config.DB.Close()
		return err
	}
	// Config.DB.Close()
	return nil
}


func GetAllProducts(prod *[]Models.Product) (err error) {
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("error in getting all user")
	// }
	
	if err = Config.DB.Find(prod).Error; err != nil {
		// Config.DB.Close()
		return err
	}
	// Config.DB.Close()
	return nil
}


func CreatProduct(prod *Models.Product,) (err error) {
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("error while creating user")
	// }
	
	temp:=Models.Product{ProductId: "PROD"+genSonyflake(), 
						ProductName: prod.ProductName, 
						Price:prod.Price,
						Quantity: prod.Quantity}
	if err = Config.DB.Create(&temp).Error; err != nil {
		// Config.DB.Close()
		return err
	}
	// Config.DB.Close()
	
	*prod=temp
	return nil
}




func UpdateUser(prod *Models.Product, id string) (err error) {
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("error while updating user")
	// }
	Config.DB.Save(prod)
	// Config.DB.Close()
	return nil
}