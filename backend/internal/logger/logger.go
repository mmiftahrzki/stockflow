package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
}

func GetLogger() *logrus.Logger {
	return logger
}
