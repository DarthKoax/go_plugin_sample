// yourPlugin.go

package main

import "plugin_test/common"

type YourPlugin struct{}

func (p *YourPlugin) PerformAction() string {
	return "Action performed by YourPlugin"
}

// func (p *YourPlugin) Auth() string {
// 	return "Authing your plugin"
// }

// Export a function that returns a new plugin instance as a common.PluginInterface.
// This is the correct way to make Plugin a function that main can look up and call.
var Plugin = func() common.PluginInterface {
	return &YourPlugin{}
}
