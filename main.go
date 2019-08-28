package main

import (
	"path/filepath"

	"progo/build/cmd/service"
	"progo/core/config"
)

func main() {
	envPath, _ := filepath.Abs("./.env")
	config.Init(envPath)

	port := ":" + config.Get("HOST_PORT")

	service.Run(port)
}
