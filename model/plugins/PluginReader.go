package plugins

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

var (
	pluginPaths []string
	pluginMap map[string]plugins.PluginI
)

type PluginReader struct {}

func init() {
	pluginPaths = getAllPlugins()
	pluginMap = make(map[string]plugins.PluginI)
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func getAllPlugins() (list []string) {
	if err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".so" {
			list = append(list, path)
		}
		return nil
	}); err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}

func (pr *PluginReader) LoadAllPlugins() error {
	for _, p := range pluginPaths {
		p, err := plugin.Open(p)
		fmt.Println(p)
		if err != nil {
			log.Fatal(err)
			return err
		}
		pName, err := p.Lookup("PluginName")
		if err != nil {
			log.Fatal(err)
			return err
		}
		newVersionManagement, _ := pInterface.(func() plugins.PluginReaderI) // assert the type of the func
		pluginMap[derefString(pName.(*string))] = newVersionManagement()
	}
	return nil
}

func (pr *PluginReader) GetOrgaInfo(credentials map[string]interface{}) (info map[string]interface{}) {
	fmt.Println("Start GetOrgaInfo in PluginReader")
	info = make(map[string]interface{})
	for k, v := range pluginMap {
		if credential, found := credentials[k]; found {
			orgaInfo, err := v.GetOrgaInfo(credential)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(orgaInfo)
			info[k] = orgaInfo
			return info
		}
	}
	fmt.Println("End GetOrgaInfo in PLuginReader")
	return nil
}

func (pr *PluginReader) GetTeamInfo() (info map[string]interface{}) {
	return nil
}

func (pr *PluginReader) GetInsightTeamInfo() (info map[string]interface{}) {
	return nil
}

func (pr *PluginReader) GetTeamRepoInfo() (info map[string]interface{}) {
	return nil
}