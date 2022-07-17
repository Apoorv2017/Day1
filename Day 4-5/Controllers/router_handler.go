package Controllers

import (
	"encoding/json"
	
	"io/ioutil"
	"net/http"
	"project/Database"
	"project/Models"
	"time"
	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context){
	id := c.Params.ByName("id")
	if id==""{
		c.String(http.StatusUnauthorized, "message:Provied valid Product Id")
		c.Abort()
		return
	}
	var prod Models.Product
	err := Database.GetProductById(&prod, id)
	if err != nil {
		c.String(http.StatusUnauthorized, "message:Product does not exist.")
		c.Abort()
		
	} else {
		c.JSON(http.StatusOK, prod)
	}
}

func GetAllProducts(c *gin.Context){
	var prod []Models.Product
	err := Database.GetAllProducts(&prod)
	if err != nil {
		c.String(http.StatusUnauthorized, "message:Problem in getting all products.")
		c.Abort()
		
	} else {
		c.JSON(http.StatusOK, gin.H{
			"product": prod,
		})
	}
}

func CreatProduct(c *gin.Context){
	var prod Models.Product
	c.BindJSON(&prod)
	if prod.ProductName=="" || prod.Price<=0 || prod.Quantity<=0 {
		c.String(http.StatusUnauthorized, "message:Please enter all the values correctly")
		c.Abort()
		return

	}
	err := Database.CreatProduct(&prod)
	if err != nil {
		c.String(http.StatusUnauthorized, "message:Some Problem occured while adding the Product.")
		c.Abort()
		
	} else {
		c.JSON(http.StatusOK, gin.H{"product_id": prod.ProductId,
									"product_name":  prod.ProductName,
									"price": prod.Price,
									"quantity": prod.Quantity,
									"message": "product successfully added",})
	}
}

func UpdateProduct(c *gin.Context){
	var prod Models.Product
	id := c.Params.ByName("id")
	if id==""{
		c.String(http.StatusUnauthorized, "message:Provied valid Product Id")
		c.Abort()
		return
	}
	err := Database.GetProductById(&prod, id)
	if err != nil {
		c.String(http.StatusUnauthorized, "message:Product was not found.")
		c.Abort()
		return
	}
	c.BindJSON(&prod)
	
	if prod.ProductName=="" || prod.Price<=0 || prod.Quantity<=0{
		c.String(http.StatusUnauthorized, "message:Provied valid Inputs")
		c.Abort()
		return
	}

	err = Database.UpdateUser(&prod, id)
	if err != nil {
		c.String(http.StatusUnauthorized, "message:Could not update product.")
		c.Abort()
	} else {
		c.JSON(http.StatusOK, prod)
	}
}

func PlaceOrder(c *gin.Context){
	
	var order Models.Order
	temp, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(temp, &order)
	if order.CustomerId==""|| order.ProductId=="" || order.Quantity<=0{
		c.String(http.StatusUnauthorized, "message:Provied valid Inputs")
		c.Abort()
		return
	}
	var Cus_order Models.Order
	err = Database.GetLastOrderByCustomer(&Cus_order, order.CustomerId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return 
	}else if time.Now().Sub(Cus_order.CreatedAt).Minutes()<5{
		c.String(http.StatusUnauthorized, "message: Customer under cool down period")
		c.Abort()
		

	}else{
		// uniqueValue := "Locked"
		// LockDuration := 300
		// if isLocked, _ := Radis.Lock(order.ProductId, uniqueValue, LockDuration); !isLocked {
		// 		c.JSON(http.StatusOK, gin.H{"ErrorMessage" : "Lock is Already Acquired"})
		// 	return 
		// }
		// // time.Sleep(10 * time.Second)
		// defer Radis.Unlock(order.ProductId, uniqueValue)
		
		var prod Models.Product
		err = Database.GetProductById(&prod, order.ProductId)
		if err != nil {
			err=Database.PlaceOrder(&order, "Failed")
			if err != nil {
				c.String(http.StatusUnauthorized, "message : Product was not found")
				c.Abort()
				return 
			}
			c.JSON(http.StatusOK, order)
		} else if order.Quantity<=prod.Quantity{
			err=Database.PlaceOrder(&order, "Order Placed")
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
			}
			prod.Quantity=prod.Quantity-order.Quantity
			err=Database.UpdateUser(&prod, prod.ProductId)
			c.JSON(http.StatusOK, order)
		}else{
			err=Database.PlaceOrder(&order, "Failed")
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
			}
			c.JSON(http.StatusOK, order)
		}
	}
}

func GetOrderById(c *gin.Context){
	id := c.Params.ByName("id")
	if id==""{
		c.String(http.StatusUnauthorized, "message:Provied valid Order ID")
		c.Abort()
		return 
	}
	var order Models.Order
	err := Database.GetOrderById(&order, id)
	if err != nil {
		c.String(http.StatusUnauthorized, "message : Order was not found")
		c.Abort()
		
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetAllOrders(c *gin.Context){
	var order []Models.Order
	err := Database.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"orders": order,
		})
	}
}