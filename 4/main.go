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
	fmt.Print("f type is %T", f)
	fmt.Print("f value is %v", f)

	fmt.Print("d type is %T", d)
	fmt.Print("d value is %v", d)
}
