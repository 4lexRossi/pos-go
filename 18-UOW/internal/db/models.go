// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import ()

type Category struct {
	ID          int32
	Name        string
}

type Course struct {
	ID          string
	CategoryID  string
	Name        int32
}