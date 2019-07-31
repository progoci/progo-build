package main

import (
	"progo/build/cmd/service"
	"progo/build/config"
)

func main() {
	config.Init()

	port := ":" + config.Get("HOST_PORT")

	service.Run(port)
}
