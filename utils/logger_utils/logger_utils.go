package logger_utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	LoggerUtils loggerUtilsInterface = &logger{}
)

type logger struct{}

type loggerUtilsInterface interface {
	InitLogger()
}

func (s *logger) InitLogger() {
	if os.Getenv("LOG_LEVEL") == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	log.Infoln("Starting application")
}
