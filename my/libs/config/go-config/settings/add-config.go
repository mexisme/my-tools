package settings

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// AddConfigItems adds a new configuration item, and makes it overridable by env vars
func AddConfigItems(configKeys []string) {
	for _, key := range configKeys {
		viper.BindEnv(key)
	}
}

// AddConfigItemsWithFlags adds a new configuration item, as above, but also binds to a PFlag
// of the same name.
// NOTE: The PFlag must already have been created with something like:
//       _ = pflag.String("host", "h", "Host to lookup")
func AddConfigItemsWithPFlags(configKeys []string) error {
	for _, key := range configKeys {
		viper.BindEnv(key)

		if err := viper.BindPFlag(key, pflag.Lookup(key)); err != nil {
			return err
		}
	}
	return nil
}
