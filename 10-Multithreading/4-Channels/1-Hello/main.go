package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() {
		c <- "hello world"
	}()
	msg := <-c
	fmt.Println(msg)
}
