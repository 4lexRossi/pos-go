package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello"

	showType(x)
	showType(y)

	var myVar interface{} = "Alex Rossi"

	println(myVar.(string))

	res, ok := myVar.(int)

	fmt.Printf("res value is %v and ok result is %v", res, ok)
}

func showType(t interface{}) {
	fmt.Printf("the variable type is %T and the value is %v\n", t, t)
}
