package config

import (
	"os"
	"poller-app/internal/service/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	logger := log.NewLogger()
	t.Run("success : successfully read config", func(t *testing.T) {
		currentDir, _ := os.Getwd()
		conf, err := ReadConfig(currentDir+"/config.json", logger)
		assert.Nil(t, err)
		assert.Equal(t, *conf.AppName, "poller-app")
	})

	t.Run("failure : config file doesn't exist in the path provided", func(t *testing.T) {
		conf, err := ReadConfig("/config.json", logger)
		assert.NotNil(t, err)
		assert.Nil(t, conf.AppName)
	})
}
