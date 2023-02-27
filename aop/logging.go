package aop

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetDefaultLogger() *log.Logger {
	logger := log.New()

	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.ErrorLevel)

	return logger
}
