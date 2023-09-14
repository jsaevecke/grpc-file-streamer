package config

import "github.com/rs/zerolog"

// ConfigLogging contains settings to control the service logging
type ConfigLogging struct {
	PrettyPrint bool          `env:"PRETTY_PRINT" envDefault:"false"`
	LogLevel    zerolog.Level `env:"LOG_LEVEL" envDefault:"info" validate:"required"`
}
