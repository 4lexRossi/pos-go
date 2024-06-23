package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)

	qtdWorkerss := 10
	for i := 0; i < qtdWorkerss; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
