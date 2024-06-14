package main

import (
	"fmt"

	"github.com/4lexRossi/pos-go/7-Packaging/4-Workspaces/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	println(uuid.New().String())
}
