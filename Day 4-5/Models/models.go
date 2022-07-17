package Models

import (
	"time"
)

type Product struct {
	ProductId string `json:"product_id" gorm:"primary_key"`
	ProductName string `json:"product_name"`
	Price int `json:"price"`
	Quantity int `json:"quantity"`
}

type Order struct {
	
	OrderId string `json:"order_id" gorm:"primary_key"`
	CustomerId string `json:"customer_id"`
	ProductId string `json:"product_id"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
	CreatedAt  time.Time `json:"-"`
}


