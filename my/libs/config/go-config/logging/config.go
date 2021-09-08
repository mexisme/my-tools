package logging

import (
	"github.com/mexisme/go-config/settings"
	errors "golang.org/x/xerrors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetFromConfig sets App name and env, Log format, debug and Sentry DSN from config settings
func (s *Config) SetFromConfig() *Config {
	settings.AddConfigItems([]string{
		configItemAppName, configItemAppEnv,
		configItemFormat, configItemSentryDsn,
	})

	s.AppName = viper.GetString(configItemAppName)
	s.AppEnv = viper.GetString(configItemAppEnv)
	// We do this before setting debug mode, to help-out the log aggregators:
	if _, err := s.SetFormat(viper.GetString(configItemFormat)); err != nil {
		log.Panic(err)
	}
	s.SetDebug(viper.GetBool(settings.ConfigItemDebug))
	s.SentryDsn = viper.GetString(configItemSentryDsn)

	return s
}

// SetDebug enables/disables debug logging
func (s *Config) SetDebug(debug bool) *Config {
	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")

	} else {
		log.SetLevel(log.InfoLevel)
	}

	return s
}

// SetFormat sets the log-output format:  either "text" or "json"
func (s *Config) SetFormat(loggingFormat string) (*Config, error) {
	switch loggingFormat {
	case "":
		fallthrough

	case "text":
		// This is the default log formatter in logrus, anyway:
		log.SetFormatter(&log.TextFormatter{})

	case "json":
		log.SetFormatter(&log.JSONFormatter{})

	default:
		return nil, errors.Errorf("Log format %#v not supported.", loggingFormat)
	}

	return s, nil
}
