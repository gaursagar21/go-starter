package main

import (
	"fmt"
	"github.com/gaursagarMT/starter/src/config"
	"github.com/gaursagarMT/starter/src/ports/grpc"
	"github.com/gaursagarMT/starter/src/utilities/logger"
	"log"
)

func main() {
	fmt.Println("Application Starting")

	// Init Config
	var conf config.Config
	conf.Get()

	// Init Loggers
	err := logger.NewLogger(logger.Configuration{
		LogLevel:   logger.DEBUG,
		EnableJSON: false,
		Output:     logger.CONSOLE,
	}, logger.LogrusLogger)

	if err != nil {
		log.Fatal(fmt.Sprintf("Could not start logger due to error: %s", err))
	}

	err = grpc.StartGRPCServer(conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not start app due to error: %s", err))
		return
	}
}
