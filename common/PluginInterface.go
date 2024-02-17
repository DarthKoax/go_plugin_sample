// common/plugininterface.go
package common

// PluginInterface defines the functions that our plugins must implement.
type PluginInterface interface {
	PerformAction() string
	// Auth() string
}
