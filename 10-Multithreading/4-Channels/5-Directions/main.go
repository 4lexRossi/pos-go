package main

import "fmt"

func main() {
	hello := make(chan string)
	go receiver("Hello", hello)
	deliver(hello)
}

func receiver(name string, hello chan<- string) {
	hello <- name
}

func deliver(data <-chan string) {
	fmt.Println(<-data)
}
