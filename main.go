// main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"plugin_test/common"
)

func main() {
	pluginsDir := "./plugins" // Directory where plugins are stored
	plugins, err := loadPlugins(pluginsDir)
	if err != nil {
		fmt.Println("Error loading plugins:", err)
		return
	}

	for _, p := range plugins {
		fmt.Println(p.PerformAction())
		// fmt.Println(p.Auth())
	}
}

func loadPlugins(dir string) ([]common.PluginInterface, error) {
	var plugins []common.PluginInterface

	// Walk through the plugin directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".so" {
			// Load the plugin
			plug, err := plugin.Open(path)
			if err != nil {
				return err
			}
			// Look up the exported symbol (an instance of PluginInterface)
			symPlugin, err := plug.Lookup("Plugin")
			if err != nil {
				return err
			}

			// Assert the correct function signature.
			newPluginFuncPtr, ok := symPlugin.(*func() common.PluginInterface)
			if !ok {
				return fmt.Errorf("unexpected type from module symbol, got type: %T", symPlugin)
			}

			// Dereference the pointer to get the function.
			// newPluginFunc := *newPluginFuncPtr

			// Use the function to create an instance of the plugin
			// pluginInstance := newPluginFunc()
			pluginInstance := (*newPluginFuncPtr)()
			plugins = append(plugins, pluginInstance)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return plugins, nil
}
