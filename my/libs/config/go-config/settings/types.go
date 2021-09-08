package settings

const (
	// ConfigItemDebug defines the Viper config item for running in debug mode
	ConfigItemDebug = "debug"
	// ConfigItemDryRun defines the Viper config item for enabling dry-run mode
	ConfigItemDryRun = "dry_run"
)

// Config provides the App Config for some of the app-wide settings
type Config struct {
	// File is the default config file name
	File string
	// Dir is an additional directory to search for config files
	Dir string
	// EnvPrefix allows you to add a Viper "EnvPrefix" to config env-vars
	EnvPrefix string
	// UseOnlyDir disables looking for a config file in "$HOME" or "." directories.
	OnlyUseDir bool
}
