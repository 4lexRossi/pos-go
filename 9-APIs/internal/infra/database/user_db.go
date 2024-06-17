package database

import "gorm.io/gorm"

type User struct {
	DB *gorm.DB
}

func newUser(db *gorm.DB) *User {
	return &User{DB: db}
}
