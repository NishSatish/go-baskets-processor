package main

import (
	"go-basket-processor/api"
	"go-basket-processor/pkg/config"
)

func main() {
	config.LoadEnv(".")

	api.StartServer()
}
