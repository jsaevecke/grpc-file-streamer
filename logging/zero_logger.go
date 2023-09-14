package logging

import (
	"os"
	"time"

	"grpc-file-streamer/apperrors"
	"grpc-file-streamer/env"

	"github.com/rs/zerolog"
)

/*
InitLogger creates a new logger instance with contextual fields to pass along to other service.
This ensures that the logger instance, when used, has the minimum required fields set and
that other services can modify the context even further through the returned logger instance.
*/
func InitLogger() (zerolog.Logger, error) {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFieldName = "@timestamp"

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	// if pretty print is enabled, format all logging messages to a more human-readable format
	if os.Getenv(env.EnvPrettyPrint) == "true" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Logger()
	}

	// set log level
	strLogLevel := os.Getenv(env.EnvLogLevel)
	if strLogLevel == "" {
		strLogLevel = "info"
	}

	logLevel, errParse := zerolog.ParseLevel(strLogLevel)
	if errParse != nil {
		return logger, apperrors.ErrValidation{
			Issue:  errParse,
			Caller: "NewLogger",
		}
	}

	logger = logger.Level(logLevel).With().Logger()

	return logger, nil
}
