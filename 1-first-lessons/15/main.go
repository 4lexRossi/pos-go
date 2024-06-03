package main

import "fmt"

type Client struct {
	name string
}

func NewClient() *Client {
	return &Client{name: "Alex"}
}

func (c *Client) walked() {
	c.name = "Alex Rossi"
	fmt.Printf("The client %v walked\n", c.name)
}

func main() {

	alex := Client{
		name: "Alex",
	}
	alex.walked()

	fmt.Println(NewClient().name)
	fmt.Printf("the struct value with name %v", alex.name)
}
