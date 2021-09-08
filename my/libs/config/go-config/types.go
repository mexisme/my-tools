package config

/*
Config provides basic fields for configuring the "settings" and "logging" packages.

"File" is the name of a file that Viper will read for configuration.
By default, it searches for the file in the user's `$HOME` dir as well as the current
workig dir -- but see "OnlyUseDir" below.
If the file-name is empty, settings won't be loaded from a file (only env-vars).

"Dir" is an optional additional additional dir to search for a config file.

"OnlyUseDir" when false will additionally search "$HOME" and current working dir for
the config file. When true, will only search in the above "Dir" directory.
If "Dir" is not given, then the config file won't be loaded.

"EnvPrefix" is a required prefix-string that Viper uses to filter Env-vars
for settings.

"Debug" enables debug logging if set to "true":

"FromConfig" enables the following settings (Name ... LoggingSentryDsn)
to be configured via Viper.
This means it will use the above Config file and appropriate Env-vars

"Name" is the App name, used in log messages.
This can be a string, or a func ref that will return a string.

"Environment" is the App's environment it was run in -- e.g. "staging" or "prod"
This can be a string, or a func ref that will return a string.

"Release" is the App's release / version string
This can be a string, or a func ref that will return a string.

"LoggingFormat" sets the log-out format for log messages

"LoggingSentryDsn" is the connection string (DSN) used to send errors to Sentry.io
*/
type Config struct {
	File       string
	Dir        string
	OnlyUseDir bool
	EnvPrefix  string
	Debug      bool

	FromConfig bool

	Name             interface{}
	Environment      interface{}
	Release          interface{}
	LoggingFormat    string
	LoggingSentryDsn string

	// We don't want to try to reinitialise the config more than once
	initConfigDone bool
	// We don't want to try to reinitialise the logging more than once
	logConfigDone bool
}

var config *Config

/////////

// Init forwards to the "config" struct's Init()
func Init(initConfig Config) {
	config = (&initConfig).Init()
}

// Init is to allow other packages to easily depend on this one,
// since most of the important logic is in init()
func (s *Config) Init() *Config {
	s.initSettingsOnce()
	s.initLoggingOnce()
	return s
}
