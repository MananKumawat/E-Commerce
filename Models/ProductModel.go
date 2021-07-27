package Models

import (
	"gorm.io/gorm"
	"sync"
)

var ProductMutex sync.Mutex
var OrderMutex sync.Mutex

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
	ProductId	uint `json:"product_id"`
	Quantity	uint   `json:"quantity"`

}