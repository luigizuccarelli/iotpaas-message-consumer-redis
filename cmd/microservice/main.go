// +build real

package main

import (
	"os"

	"github.com/luigizuccarelli/iotpaas-message-consumer/pkg/connectors"
	"github.com/luigizuccarelli/iotpaas-message-consumer/pkg/handlers"
	"github.com/luigizuccarelli/iotpaas-message-consumer/pkg/validator"
	"github.com/microlib/simple"
)

var (
	logger *simple.Logger
)

// Main function : keep things clean and simple
// Allows for simple E2E testing and code coverage
func main() {

	if os.Getenv("LOG_LEVEL") == "" {
		logger = &simple.Logger{Level: "info"}
	} else {
		logger = &simple.Logger{Level: os.Getenv("LOG_LEVEL")}
	}
	err := validator.ValidateEnvars(logger)
	if err != nil {
		os.Exit(-1)
	}
	conn := connectors.NewClientConnectors(logger)
	handlers.Init(conn)
	//defer conn.Close()
}
