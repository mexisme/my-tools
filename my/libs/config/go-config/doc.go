/*
Package "config" provides an abstraction of a common pattern I use for configuring my application.
This includes configuring logging (via the "Logrus" package) so I can log debug-level information
around configuration at the same time.

The "config" package is a simple abstraction around the "settings" and "logging" sub-packages.

Initialisation

You're expected to initalise this by calling the Init() function with a Config{}
struct (defined below).  The struct needs to have values set in it for configuring the above
libraries.  Alternatively, you can enable the "FromConfig" setting, and it will
try to self-configure via the Viper script.
*/

package config
