package main

import "fmt"

func myPanic() {
	panic("xpto")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover panic: ", r)
		}
	}()
	myPanic()
}
