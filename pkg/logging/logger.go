package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(environment string) *logrus.Logger {
	// Initialize logger
	logger := logrus.New()

	// Set log level
	switch environment {
	case "local":
		logger.SetLevel(logrus.DebugLevel)
	case "dev":
		logger.SetLevel(logrus.InfoLevel)
	case "prod":
		logger.SetLevel(logrus.WarnLevel)
	default:
		logger.SetLevel(logrus.WarnLevel)
	}

	// Set log format
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log output
	logger.SetOutput(os.Stdout)

	return logger
}
