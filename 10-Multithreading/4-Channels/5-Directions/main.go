package main

import "fmt"

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	deliver(hello)
}

func receive(name string, hello chan<- string) {
	hello <- name
}

func deliver(data <-chan string) {
	fmt.Println(<-data)
}
