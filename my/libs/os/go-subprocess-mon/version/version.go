package version

const (
	application = "subprocess-monitor"
	release     = "0.2.2"
)

// Application is the "friendly" name for this code
func Application() string {
	return application
}

// Release is the current version of "ucms-smarties"
func Release() string {
	return release
}
