package main

import (
	"context"
	"database/sql"

	"github.com/4lexRossi/pos-go/17-SQLC/internal/db"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "back",
		Description: sql.NullString{String: "Desc", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "1a531a62-5292-4fa7-a258-bca6e7c2bfb7",
	// 	Name:        "Backend updated",
	// 	Description: sql.NullString{String: "Backend description updated", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.DeleteCategory(ctx, "1e307590-9774-4708-b843-c12b2b2a0a3a")
	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
