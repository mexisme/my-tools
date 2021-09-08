package config

import (
	"time"

	"github.com/spf13/viper"

	"github.com/mexisme/go-config/settings"
)

// AddConfigItems passes the configItems through to settings.AddConfigItems()
func AddConfigItems(configKeys []string) {
	settings.AddConfigItems(configKeys)
}

// AddConfigItemsWithFlags passes the configItems through to settings.AddConfigItemsWithFlags()
func AddConfigItemsWithPFlags(configKeys []string) error {
	return settings.AddConfigItemsWithPFlags(configKeys)
}

// Get forwards to viper.Get
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetBool forwards to viper.GetBool
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetFloat64 forwards to viper.GetFloat64
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetInt forwards to viper.GetInt
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetString forwards to viper.GetString
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStringMap forwards to viper.GetStringMap
func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

// GetStringMapString forwards to viper.GetStringMapString
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// GetStringSlice forwards to viper.GetStringSlice
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetTime forwards to viper.GetTime
func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

// GetDuration forwards to viper.GetDuration
func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

// IsSet forwards to viper.IsSet
func IsSet(key string) bool {
	return viper.IsSet(key)
}

// AllSettings forwards to viper.AllSettings
func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}
