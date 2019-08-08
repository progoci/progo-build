package log

import (
	logger "log"

	"progo/build/config"
)

var levels = map[string]int{
	"fatal": 50,
	"error": 40,
	"warn":  30,
	"info":  20,
	"debug": 10,
}

var configuredLevel int

// Print outputs the given values if the logging level agrees.
func Print(level string, message string, v ...interface{}) {

	// The level of the given values to log.
	logLevel, ok := levels[level]
	if !ok {
		logger.Print("Invalid log level" + level)
		return
	}

	configured := getConfiguredLevel()

	// Logs value if its level is at least at the configured level.
	if logLevel >= configured {
		if message != "" {
			logger.Print(message)
		}

		logger.Print(v...)
	}
}

func getConfiguredLevel() int {

	if configuredLevel != 0 {
		return configuredLevel
	}

	levelConfig := config.Get("LOGGING_LEVEL")
	if levelConfig == "" {
		levelConfig = "error"
	}

	// The configured log level.
	var ok bool
	if configuredLevel, ok = levels[levelConfig]; !ok {
		configuredLevel = levels["error"]
	}

	return configuredLevel
}
