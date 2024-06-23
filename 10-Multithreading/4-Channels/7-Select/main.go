package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{Msg: "kafka", ID: 1}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{Msg: "RabbitMQ", ID: 2}
			c2 <- msg
		}
	}()

	// infinit for
	for {
		// for i := 0; i < 3; i++ {
		select {
		case msg := <-c1:
			fmt.Printf("received Kafka: %s/n", msg.Msg)
		case msg := <-c2:
			fmt.Printf("received RabbitMQ: %s/n", msg.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
