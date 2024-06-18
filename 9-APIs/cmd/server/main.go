package main

import (
	"net/http"

	"github.com/4lexRossi/pos-go/9-APIs/internal/entity"
	"github.com/4lexRossi/pos-go/9-APIs/internal/infra/database"
	"github.com/4lexRossi/pos-go/9-APIs/internal/infra/webservers/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// _, err := configs.LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products{id}", productHandler.GetProduct)
	r.Put("/products{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":8000", r)
}
