package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/mexisme/go-config/logging"
	"github.com/mexisme/go-config/settings"
)

func (s *Config) initSettingsOnce() {
	// This should make it safe to rerun a few times
	if !s.initConfigDone {
		if err := (&settings.Config{
			File:       s.File,
			Dir:        s.Dir,
			EnvPrefix:  s.EnvPrefix,
			OnlyUseDir: s.OnlyUseDir,
		}).Read(); err != nil {
			log.Panic(err)
		}

		s.initConfigDone = true
	}
}

func (s *Config) initLoggingOnce() {
	// We need to make sure the settings have been read at least once:
	s.initSettingsOnce()

	// This should make it safe to rerun a few times
	if !s.logConfigDone {
		logConfig := &logging.Config{
			AppName:    MissingValueIsEmpty(FromStringOrFunc(s.Name)).(string),
			AppEnv:     MissingValueIsEmpty(FromStringOrFunc(s.Environment)).(string),
			AppRelease: MissingValueIsEmpty(FromStringOrFunc(s.Release)).(string),
			SentryDsn:  MissingValueIsEmpty(FromStringOrFunc(s.LoggingSentryDsn)).(string),
		}

		logConfig.SetFormat(s.LoggingFormat)

		if s.FromConfig {
			logConfig.SetFromConfig()
		}

		logConfig.Init()

		s.logConfigDone = true
	}
}
