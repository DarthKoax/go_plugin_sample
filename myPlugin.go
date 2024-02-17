// myPlugin.go

package main

import "plugin_test/common"

type MyPlugin struct{}

func (p *MyPlugin) PerformAction() string {
	return "Action performed by MyPlugin"
}

// Export a function that returns a new plugin instance as a common.PluginInterface.
// This is the correct way to make Plugin a function that main can look up and call.
var Plugin = func() common.PluginInterface {
	return &MyPlugin{}
}
