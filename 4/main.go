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
	fmt.Printf("f type is %T", f)
	fmt.Printf("f value is %v", f)

	fmt.Printf("d type is %T", d)
	fmt.Printf("d value is %v", d)
}
