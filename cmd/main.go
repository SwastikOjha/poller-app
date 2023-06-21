package main

import (
	"fmt"
	"log"
	"poller-app/cmd/config"
	"poller-app/internal/base"
	"poller-app/internal/engine/apipoller"
	logger "poller-app/internal/service/log"
)

func main() {
	logger := logger.NewLogger()
	fmt.Println("app started")
	config, err := config.ReadConfig(base.ConfigPath, logger)
	if err != nil {
		logger.Println(err)
	}
	err = runMain(logger, config)
	if err != nil {
		logger.Println(err)
	}
	fmt.Print(&base.Buf)
}

func runMain(logger *log.Logger, config config.Config) error {
	ap := apipoller.NewApiPoller(logger, config)
	err := ap.Run()
	if err != nil {
		logger.Printf("%v", err)
		return err
	}
	return nil
}
