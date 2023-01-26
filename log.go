package goutils

import (
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/lokirus"
)

var (
	logger *logrus.Logger
)

func getLokiHook() *lokirus.LokiHook {

	// Configure the Loki hook
	opts := lokirus.NewLokiHookOptions().
		// Grafana doesn't have a "panic" level, but it does have a "critical" level
		// https://grafana.com/docs/grafana/latest/explore/logs-integration/
		WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
		WithFormatter(&logrus.JSONFormatter{}).
		WithStaticLabels(lokirus.Labels{
			"app":         GetEnvVariable("APP_NAME", "unset"),
			"environment": GetEnvVariable("ENVIRONMENT", "development"),
		}).
		WithBasicAuth(GetEnvVariable("LOKI_USER", ""), GetEnvVariable("LOKI_PASS", "")) // Optional

	return lokirus.NewLokiHookWithOpts(
		GetEnvVariable("LOKI_URL", ""),
		opts,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel)
}

// Logger returns a new logger
func Logger() *logrus.Logger {
	if logger == nil {

		log := logrus.New()

		logLevel, err := logrus.ParseLevel(GetEnvVariable("LOG_LEVEL", "info"))
		if err == nil {
			log.SetLevel(logLevel)
		}

		if GetEnvVariable("LOKI_URL", "") != "" {
			log.AddHook(getLokiHook())
		}

		logger = log

	}

	return logger
}
