package main

import (
	"fmt"
	"pos-go/math"

	"github.com/google/uuid"
)

func main() {
	s := math.Sum(10, 20)
	fmt.Println("The result is", s)
	fmt.Println(uuid.New())
}
