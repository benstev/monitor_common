package initializers

import (
	// config "gomonitor/internal/services/config"

	"os"
	"strconv"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// LogLevelEnv is an environment variable name for LOG_LEVEL
	LogLevelEnv = "LOG_LEVEL"
	// EnableJSONLogsEnv is an environment variable name for ENABLE_JSON_LOGS
	EnableJSONLogsEnv = "ENABLE_JSON_LOGS"
	// DefaultLogLevel is a default LOG_LEVEL value
	DefaultLogLevel = 1
)

// InitializeLogs setups zerolog logger
func InitializeLogs() error {
	logLevel := DefaultLogLevel
	if logLevelStr := envy.Get(LogLevelEnv, ""); logLevelStr != "" {
		logLevel, _ = strconv.Atoi(logLevelStr)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	// json or human readable output
	if envy.Get(EnableJSONLogsEnv, "false") == "false" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	// add filepath+row num to log(app/cmd/serve.go:37)
	log.Logger = log.With().Caller().Logger()

	return nil
}
