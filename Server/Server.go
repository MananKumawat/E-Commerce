package Server

import (
	"github.com/MananKumawat/E-Commerce/Models"
	"fmt"
	"time"
)

const Capacity = 100

var(
	OrderChannel = make(chan *Models.Order, Capacity)
	PendingChannel = make(chan *Models.Order,Capacity)
)

func Start(){
	for count := 0; count < Capacity; count++ {
		go HandleOrder()
	}
	go HandlePendingOrder()
}

func HandleOrder(){
	var order *Models.Order
	for {
		order = <- OrderChannel
		var flag0 = true
		var flag1 = true
		var gottime time.Duration
		var customerorders []Models.Order

		customer_id := order.CustomerId
		Models.GetOrderByCustomerID(&customerorders, customer_id)

		for _, pastorder := range customerorders{
			if pastorder.Status == "placed" {
				gottime = time.Now().Sub(pastorder.CreatedAt)
				if gottime.Minutes() < 2 {
					flag0 = false
					break
				}
			}
		}

		for _, item := range order.Items {
			var product Models.Product
			gg := item.ProductId
			Models.GetProductByID(&product, gg)
			if product.Quantity < item.Quantity {
				flag1 = false
				break
			}
		}

		if flag0 && flag1 {
			for _, item := range order.Items {
				var product Models.Product
				Models.GetProductByID(&product, item.ProductId)
				product.Quantity -= item.Quantity
				Models.UpdateProduct(&product)
			}
			order.Status = "placed"
			Models.CreateOrder(order)
			fmt.Println("Order ",order.ID,"Placed")
		} else {
			if flag1 == false {
				order.Status = "failed"
				fmt.Println("Order ",order.ID,"failed Item quantity insufficient")
			} else {
				order.Status = "Pending"
				fmt.Println("Order ",order.ID,"is now pending")
				PendingChannel<-order
			}
		}
	}

}

func HandlePendingOrder(){
	var order *Models.Order
	for {
		order = <- PendingChannel
		gottime := time.Now().Sub(order.CreatedAt)
		if gottime.Minutes() >= 2 {
			OrderChannel<-order
		} else {
			PendingChannel<-order
		}
		time.Sleep(time.Millisecond * 100)
	}
}
