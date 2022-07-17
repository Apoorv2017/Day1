package Database

import (
	"project/Config"
	"project/Models"
	"errors"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
  	"github.com/sony/sonyflake"
	
	"strconv"
)


func genSonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	
	return strconv.FormatUint(id, 10)
}

func PlaceOrder(order *Models.Order, status string) (err error){
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("Database not accessable")
	// }

	
	temp:=Models.Order{OrderId: "ORD"+genSonyflake(), 
						CustomerId: order.CustomerId,
						ProductId: order.ProductId, 
						Quantity: order.Quantity,
						Status:status}
	if status== "failed"{
		if err = Config.DB.Create(&temp).Error; err != nil {
			// Config.DB.Close()
			return err
		}
		// Config.DB.Close()
		return nil
	}else{
		if err = Config.DB.Create(&temp).Error; err != nil {
			// Config.DB.Close()
			return err
		}
		// Config.DB.Close()
		
		*order=temp
		return nil
	}
}

func GetOrderById(order *Models.Order, id string) (err error) {
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("error while getting user by id")
	// }
	
	if err = Config.DB.Where("order_id= ?", id).First(order).Error; err != nil {
		// Config.DB.Close()
		return err
	}
	// Config.DB.Close()
	return nil
}

func GetAllOrders(order *[]Models.Order) (err error) {
	// db, err := gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("error in getting all user")
	// }
	if err = Config.DB.Find(order).Error; err != nil {
		// Config.DB.Close()
		return err
	}
	// Config.DB.Close()
	return nil
}

func GetLastOrderByCustomer(order *Models.Order, id string)(err error){
	// db, err:=gorm.Open("sqlite3", "Database.db")
	// if err!=nil{
	// 	panic("Database could not be opened.")
	// }
	err=Config.DB.Where("customer_id = ?", id).Last(&order).Error
	if  errors.Is(err, gorm.ErrRecordNotFound) {
		// Config.DB.Close()
		return nil
	}else if err!=nil{
		// Config.DB.Close()
		return err
	}else{
		// Config.DB.Close()
		return nil
	}
}