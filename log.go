package goutils

import (
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

//Logger returns a new logger
func Logger() *logrus.Logger {
	if logger == nil {

		log := logrus.New()

		logLevel, err := logrus.ParseLevel(GetEnvVariable("LOG_LEVEL", "info"))
		if err == nil {
			log.SetLevel(logLevel)
		}

		logger = log

	}

	return logger
}
