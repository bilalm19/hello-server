package main

import (
	"hello-server/internal/logger"
	"hello-server/internal/server"
)

func main() {
	helloServer := server.New()
	logger.Logger.Info("Starting server")
	if err := helloServer.StartHelloServer(); err != nil {
		logger.Logger.Fatal(err)
	}
}
