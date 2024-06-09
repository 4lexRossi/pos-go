package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/lexlabs"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create one
	// db.Create(&Product{
	// 	Name:  "Laptop",
	// 	Price: 1000.00,
	// })

	//create batch
	products := []Product{
		{Name: "Mouse", Price: 200.15},
		{Name: "Mouse Pad", Price: 20.48},
		{Name: "Keyboard", Price: 500.95},
	}
	db.Create(&products)
}
