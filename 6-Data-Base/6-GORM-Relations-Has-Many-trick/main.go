package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/lexlabs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	//create product
	db.Create(&Product{
		Name:       "Laptop",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	//create serial number
	db.Create(&SerialNumber{
		Number:    "1234",
		ProductID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	//search has many
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, "/", category.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
