package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("first line")
	fmt.Println("second line")
	fmt.Println("third line")
}
