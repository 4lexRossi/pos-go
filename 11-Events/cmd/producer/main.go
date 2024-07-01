package main

import "github.com/4lexRossi/pos-go/11-Events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello RabbitMQ", "amq.direct")
}
