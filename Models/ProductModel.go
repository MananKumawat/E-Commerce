package Models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name    	string `json:"name"`
	Price 		uint   `json:"price"`
	Quantity	uint   `json:"quantity"`
}

type Order struct {
	gorm.Model
	Status 		string `json:"status"`
	CustomerId 	string `json:"customer_id"`
	Items 		[]Item
}

type Item struct {
	gorm.Model
	OrderId		uint
	Name 		string `json:"name"`
	Price 		uint   `json:"price"`
	Quantity	uint   `json:"quantity"`
}