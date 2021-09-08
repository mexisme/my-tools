package settings

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	errors "golang.org/x/xerrors"
)

// Read uses Viper to read the configuration from .config.* files or Env Vars
func (s *Config) Read() error {
	// This means any "." chars in a FQ config name will be replaced with "_"
	// e.g. "sentry.dsn" --> "$SENTRY_DSN" instead of "$SENTRY.DSN" (which won't work)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if s.EnvPrefix != "" {
		viper.SetEnvPrefix(s.EnvPrefix)
	}
	viper.BindEnv(ConfigItemDebug)
	viper.BindEnv(ConfigItemDryRun)

	if s.File != "" {
		if s.OnlyUseDir && s.Dir == "" {
			return errors.Errorf("Provided 'Dir' is empty, but 'OnlyUseDir' is set?")
		}

		viper.SetConfigName(s.File)

		// Set the config dir's to search for the given file-name.
		// It selects the first one it finds:
		if s.Dir != "" {
			viper.AddConfigPath(s.Dir)
		}

		if !s.OnlyUseDir {
			viper.AddConfigPath("$HOME")
			viper.AddConfigPath(".")
		}

		if err := viper.ReadInConfig(); err == nil {
			log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Debug("Using file")

		} else {
			var errorNotFound viper.ConfigFileNotFoundError
			if !errors.As(err, &errorNotFound) {
				log.Errorf("Error: %#v", err)
				return errors.Errorf("%w, when reading %#v", err, viper.ConfigFileUsed())
			}

			// It might be valid for there to not-be a config file
			// e.g. when using env-vars for config
			log.WithField("config_file", viper.ConfigFileUsed()).Infof("Skipping reading config: %v", err)
		}
	}

	return nil
}
