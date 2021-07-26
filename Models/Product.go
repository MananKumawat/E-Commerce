package Models

import (
	"awesomeProject/Day45/Exercise/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)
//GetAllProducts Fetch all product data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Preload(clause.Associations).Find(product).Error; err != nil {
		return err
	}
	return nil
}
//CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID ... Fetch only one product by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
//UpdateProduct ... Update product
func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}
//DeleteProduct ... Delete product
func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}