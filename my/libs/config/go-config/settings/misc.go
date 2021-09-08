package settings

import (
	"github.com/spf13/viper"
)

// DryRun says whether the dry_run config has been set
func DryRun() bool {
	// Note: Not being set should count as "false"
	return viper.GetBool(ConfigItemDryRun)
}
