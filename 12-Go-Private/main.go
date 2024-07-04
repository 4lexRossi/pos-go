package main

import (
	"fmt"

	"github.com/4lexRossi/llutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
