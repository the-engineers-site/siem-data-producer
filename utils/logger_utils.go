package utils

import log "github.com/sirupsen/logrus"

var (
	LoggerUtils loggerUtilsInterface = &logger{}
)

type logger struct{}

type loggerUtilsInterface interface {
	InitLogger()
}

func (s *logger) InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Infoln("Starting application")
}
