package main

import "fmt"

type Address struct {
	Street string
	number int
	City   string
	State  string
}

type Person interface {
	Deactivate()
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Client) Deactivate() {
	c.Active = false
}

func Deactivated(person Person) {
	person.Deactivate()
}

func main() {
	alex := Client{
		Name:   "Alex",
		Age:    38,
		Active: true,
	}
	fmt.Println(alex)

	alex.number = 10
	alex.City = "Piracity"
	alex.Deactivate()

	fmt.Println(alex)
}
