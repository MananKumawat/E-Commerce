package main

import (
	"awesomeProject/Day45/Exercise/Config"
	"awesomeProject/Day45/Exercise/Models"
	"awesomeProject/Day45/Exercise/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

)
var err error
func main() {

	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.Product{})

	r := Routes.SetupRouter()
	//running
	r.Run()

}