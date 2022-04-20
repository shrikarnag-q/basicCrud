package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Change DB details as per your msql server
const DNS = "root:admin123@tcp(127.0.0.1:3306)/product_crud?charset=utf8mb4&parseTime=True&loc=Local"

func InitializeDB() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("Not connecting to DB ", err.Error())
		panic("Not connecting.. Oh My God")
	}
	fmt.Println("Connected to DB")
	// Migrate the schema
	DB.AutoMigrate(&Product{})

}
