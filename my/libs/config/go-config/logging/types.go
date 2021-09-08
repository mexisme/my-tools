package logging

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	configItemAppName   = "application.name"
	configItemAppEnv    = "application.environment"
	configItemFormat    = "logging.format"
	configItemSentryDsn = "logging.sentry.dsn"
)

// Config contains the details for configuring Logging
type Config struct {
	AppName    string
	AppEnv     string
	AppRelease string
	SentryDsn  string
}

// Logging is a singleton for managing logging
var Logging *Config

// New creates a new struct for managing logging
func New() *Config {
	return &Config{}
}

// Init set-ups logging for the Logrus library, incl. App name, environment and release, and Sentry DSN
func (s *Config) Init() {
	appName := s.AppName
	appRelease := s.AppRelease
	appEnv := s.AppEnv

	if appName == "" {
		appName = os.Args[0]
	}
	if appRelease != "" {
		appRelease = fmt.Sprintf(" release %v", appRelease)
	}
	if appEnv != "" {
		appEnv = fmt.Sprintf(" (in %v)", appEnv)
	}

	// TODO: Should this be a Debug message?
	log.Infof("## %#v%v%v ##", appName, appRelease, appEnv)

	if s.SentryDsn != "" {
		if err := s.setupSentry(s.SentryDsn); err != nil {
			log.Error(err)
		}
	}
}
