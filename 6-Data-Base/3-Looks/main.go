package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/lexlabs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	//create category
	// category := Category{Name: "Eletronics"}
	// db.Create(&category)

	// category2 := Category{Name: "kitchen"}
	// db.Create(&category2)

	// //create product
	// db.Create(&Product{
	// 	Name:       "Laptop",
	// 	Price:      1000.00,
	// 	Categories: []Category{category, category2},
	// })

	tx := db.Begin()
	var c Category

	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronics"
	tx.Debug().Save(&c)
	tx.Commit()
}
