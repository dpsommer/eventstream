package logging

import (
	"log"
	"os"
)

func NewLogger(prefix string) *log.Logger {
	logger := log.New(os.Stdout, prefix, log.LstdFlags)
	logger.SetFlags(log.LstdFlags | log.Lshortfile)

	return logger
}
