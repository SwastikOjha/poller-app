package filehandler

import (
	"poller-app/internal/domain"
	"poller-app/internal/service/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	logger := log.NewLogger()
	t.Run("success: file is present", func(t *testing.T) {
		fileName := "../../../cmd/config/config.json"
		val, err := ReadFile(fileName, logger)
		assert.Nil(t, err)
		assert.NotNil(t, val)
	})
	t.Run("failure: file is missing", func(t *testing.T) {
		fileName := ""
		val, err := ReadFile(fileName, logger)
		assert.NotNil(t, err)
		assert.Nil(t, val)
	})
}

func TestWriteFile(t *testing.T) {
	logger, result := log.NewLogger(), "success"
	t.Run("success: write success", func(t *testing.T) {
		err := WriteFile(domain.Response{Result: &result}, "/tmp/testwritefile.json", logger)
		assert.Nil(t, err)
	})
	t.Run("failure: write success", func(t *testing.T) {
		err := WriteFile(domain.Response{Result: &result}, "", logger)
		assert.NotNil(t, err)
	})

}
