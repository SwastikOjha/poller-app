package apipoller

import (
	"fmt"
	"os"
	"poller-app/cmd/config"
	"poller-app/internal/service/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunLogic(t *testing.T) {
	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	log := log.NewLogger()
	conf, err := config.ReadConfig("../../../cmd/config/config.json", log)
	assert.Nil(t, err)
	ap := ApiPoller{Logger: log, Config: conf}
	err = runLogic(ap)
	assert.Nil(t, err)
}
