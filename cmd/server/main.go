package main

import "github.com/pcechetto/go-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
