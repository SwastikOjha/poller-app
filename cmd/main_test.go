package main

import (
	"fmt"
	"os"
	"poller-app/cmd/config"
	logger "poller-app/internal/service/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunMain(t *testing.T) {
	logger := logger.NewLogger()
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	config, err := config.ReadConfig(currentDir+"/config/config.json", logger)
	if err != nil {
		logger.Println(err)
	}
	if err != nil {
		panic(err)
	}
	apiKey := *config.ApiKey
	t.Run("failure api key missing", func(t *testing.T) {
		*config.ApiKey = ""
		assert.NotNil(t, runMain(logger, config))
	})
	t.Run("failure wrong currency denomination", func(t *testing.T) {
		*config.ApiKey = apiKey
		*config.BaseCurrency = "kjhkjhkjh"
		fmt.Println(runMain(logger, config))
		assert.NotNil(t, runMain(logger, config))
	})
	t.Run("failure invalid target file path", func(t *testing.T) {
		*config.ApiKey = apiKey
		*config.BaseCurrency = "INR"
		*config.TargetFilePath = ""
		fmt.Println(runMain(logger, config))
		assert.NotNil(t, runMain(logger, config))
	})
}
