package main

import (
	"github.com/brunobandev/reppyapi/cmd/http/handler"
	"github.com/brunobandev/reppyapi/internal/container"
)

func main() {
	c := container.Inject()
	handler.StartServer(c)
}
