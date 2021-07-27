package Controllers

import (
	"github.com/MananKumawat/E-Commerce/Models"
	"github.com/MananKumawat/E-Commerce/Server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
//GetOrders ... Get all orders
func GetOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
//CreateOrder ... Create order
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	order.Status = ""
	Server.OrderChannel <- &order
	time.Sleep(time.Second * 5)
	c.JSON(http.StatusOK, order)
	}

//GetOrderByID ... Get the order by id
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
//GetOrderByCustomerID ... Get the order by id
func GetOrderByCustomerID(c *gin.Context) {
	customer_id := c.Params.ByName("customer_id")
	var order []Models.Order
	err := Models.GetOrderByCustomerID(&order, customer_id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//UpdateOrder ... Update the order information
func UpdateOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	c.BindJSON(&order)
	err = Models.UpdateOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
//DeleteOrder ... Delete the order
func DeleteOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.DeleteOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

//GetOrderPlaced ... Get all placed orders
func GetOrderPlaced(c *gin.Context){
	var order []Models.Order
	err := Models.GetOrderPlaced(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}