package log

import (
	"log"
	"poller-app/internal/base"
)

func NewLogger() *log.Logger {
	return log.New(&base.Buf, "logger: ", log.Lshortfile)
}
