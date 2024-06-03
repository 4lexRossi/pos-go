package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 4, 53, 123, 4, 322, 21))
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
