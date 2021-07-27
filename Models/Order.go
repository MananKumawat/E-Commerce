package Models

import (
	"awesomeProject/Day45/Exercise/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)
//GetAllOrders Fetch all order data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Preload(clause.Associations).Find(order).Error; err != nil {
		return err
	}
	return nil
}
//CreateOrder ... Insert New data
func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderByID ... Fetch only one order by Id
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderByCustomerID ... Fetch only one order by Id
func GetOrderByCustomerID(order *[]Order, customer_id string) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("customer_id = ?", customer_id).Find(order).Error; err != nil {
		return err
	}
	return nil
}
//UpdateOrder ... Update order
func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}
//DeleteOrder ... Delete order
func DeleteOrder(order *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(order)
	return nil
}
//GetOrderPlaced ... Get all placed order
func GetOrderPlaced(order *[]Order) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("status = ?","placed").Find(order).Error; err != nil {
		return err
	}
	return nil
}