package main

import (
	"API/cmd/server"
	"API/pkg/config"
)

func main() {
	config.LoadAllConfigs()
	server.Serve()
}
