package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// Configuration represents application configuration.
type Configuration struct {
	LogLevel   string `envconfig:"LOG_LEVEL" required:"true"`
	LogFormat  string `envconfig:"LOG_FORMAT" required:"true"`
	PGDatabase string `envconfig:"POSTGRES_DATABASE" required:"true"`
	PGHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	PGParams   string `envconfig:"POSTGRES_PARAMS"`
	PGPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PGPort     int    `envconfig:"POSTGRES_PORT" required:"true"`
	PGUsername string `envconfig:"POSTGRES_USERNAME" required:"true"`
	Port       int    `envconfig:"PORT" required:"true"`
}

// MustConfig returns configuration populated from environment variables.
func MustConfig() *Configuration {
	cfg := &Configuration{}
	envconfig.MustProcess("", cfg)
	return cfg
}

// MustLogger creates application logger.
func MustLogger(level, format string) *logrus.Logger {
	var log = logrus.New()

	var logLevel logrus.Level
	switch level {
	case "debug":
		logLevel = logrus.DebugLevel
	case "info":
		logLevel = logrus.InfoLevel
	case "warn":
		logLevel = logrus.WarnLevel
	case "error":
		logLevel = logrus.ErrorLevel
	default:
		panic(fmt.Sprintf("Unknown log level: %s", level))
	}
	log.Level = logLevel

	var logFormatter logrus.Formatter
	switch format {
	case "text":
		logFormatter = &logrus.TextFormatter{}
	case "json":
		logFormatter = &logrus.JSONFormatter{}
	default:
		panic(fmt.Sprintf("Unknown log format: %s", format))
	}
	log.Formatter = logFormatter

	return log
}
