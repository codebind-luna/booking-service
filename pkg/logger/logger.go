package logger

import (
	"os"
	"strings"
	"time"

	"github.com/codebind-luna/booking-service/internal/constants"
	log "github.com/sirupsen/logrus"
)

func setTextFormatter() *log.TextFormatter {
	formatter := &log.TextFormatter{FullTimestamp: true}
	return formatter
}

func setJSONFormatter() *log.JSONFormatter {
	formatter := &log.JSONFormatter{TimestampFormat: time.RFC3339}
	formatter.DisableTimestamp = false
	return formatter
}

func ConfigureLogging() *log.Logger {
	logger := log.New()
	format, ok := os.LookupEnv(constants.EnvLogFormat)
	if !ok {
		logger.SetFormatter(setTextFormatter())
	} else {
		switch strings.ToLower(strings.Trim(format, "")) {
		case "json":
			logger.SetFormatter(setJSONFormatter())
		case "text":
			logger.SetFormatter(setTextFormatter())
		}
	}

	// Set log level
	logLevel, _ := os.LookupEnv(constants.EnvLogLevel)

	level, _ := log.ParseLevel(logLevel)

	logger.SetLevel(level)
	return logger
}
