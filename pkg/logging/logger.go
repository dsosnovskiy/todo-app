package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(environment string) *logrus.Logger {
	logger := logrus.New()

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

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		FullTimestamp:          true,
	})

	logger.SetOutput(os.Stdout)

	return logger
}
