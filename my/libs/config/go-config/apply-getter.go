package config

import (
	"github.com/spf13/viper"
)

// ApplyWith gets a setting from viper, and passes it to a closure
func ApplyWith(key string, f func(interface{})) {
	if viper.IsSet(key) {
		f(viper.Get(key))
	}
}
