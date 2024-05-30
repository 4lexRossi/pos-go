package main

import "fmt"

const a = "hello, world!"

type ID int

var (
	b bool    = true
	i int     = 10
	s string  = "Alex"
	f float64 = 1.2
	d ID      = 1
)

func main() {
	var myArray [3]int

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println(len(myArray) - 1)
	fmt.Println(len(myArray))
	fmt.Println((myArray[len(myArray)-1]))

	for i, v := range myArray {
		fmt.Printf("index value is %d and it value is %d\n", i, v)
	}
}
