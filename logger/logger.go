package logger

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func initLogger() {
	godotenv.Load(".env")

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	log = logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Warn("Invalid log level, defaulting to info level")
		level = logrus.InfoLevel
	}

	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})
	log.SetLevel(level)
}

func GetLogger() *logrus.Logger {
	if log == nil {
		initLogger()
	}
	return log
}
