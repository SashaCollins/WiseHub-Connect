/*
Loads the plugins based on the .so files located in the respective folder within the plugins folder
 */
package viewmodel

import (
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

var (
	pluginPaths []string
)

type PluginLoader struct {}

func init() {
	pluginPaths = getAllPluginPaths()
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
/*
Loads the absolute path from plugins
 */
func getAllPluginPaths() (list []string) {
	if err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".so" {
			list = append(list, path)
		}
		return nil
	}); err != nil {
		log.Printf("walk error [%v]\n", err)
	}
	return list
}
/*
Loads the Plugins, based on the pluginPaths, with all functions and methods as a shared object to be used
 */
func (pr *PluginLoader) LoadAllPlugins() (map[string]plugins.PluginI, error) {
	pluginMap := make(map[string]plugins.PluginI)
	for _, p := range pluginPaths {
		p, err := plugin.Open(p)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		pName, err := p.Lookup("PluginName")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		pInterface, err := p.Lookup("NewPlugin")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		newPlugin, _ := pInterface.(func() plugins.PluginI) // assert the type of the func
		pluginMap[derefString(pName.(*string))] = newPlugin()
	}
	return pluginMap, nil
}