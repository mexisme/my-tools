package logging

import (
	"time"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (s *Config) setupSentry(sentryDsn string) error {
	log.WithFields(log.Fields{"sentry.dsn": sentryDsn}).Debug("Configuring connection to Sentry.io")

	// TODO: Meta-tag for environment
	// Some meta-tags
	tags := map[string]string{
		// TODO: Pick a better name, as this maps to "${CONFIG}_APP":
		"app": viper.GetString("app"),
	}

	// Sentry will only log for messages of the following severity:
	levels := []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	}

	// Hook Sentry into Logrus:
	hook, err := logrus_sentry.NewWithTagsSentryHook(sentryDsn, tags, levels)
	if err != nil {
		return err
	}

	if s.AppRelease != "" {
		// Set the Sentry "release" version:
		log.WithFields(log.Fields{"release": s.AppRelease}).Debug("Setting release version in Sentry")
		hook.SetRelease(s.AppRelease)
	}
	if s.AppEnv != "" {
		// Set the Sentry "environment":
		log.WithFields(log.Fields{"environment": s.AppEnv}).Debug("Setting environment in Sentry")
		hook.SetEnvironment(s.AppEnv)
	}

	hook.StacktraceConfiguration.Enable = true

	// It seems as if the default 100ms is too short:
	hook.Timeout = 1 * time.Second

	// Now, add it into the Logrus hook-chain
	log.AddHook(hook)

	log.Info("Sentry enabled")

	return nil
}
