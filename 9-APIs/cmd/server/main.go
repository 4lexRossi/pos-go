package main

import "github.com/4lexRossi/pos-go/9-APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
