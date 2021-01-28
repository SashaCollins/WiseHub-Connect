package plugins

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/model/plugins/continuous_integration"
	"github/SashaCollins/Wisehub-Connect/model/plugins/version_management"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

var (
	testingToolPluginPaths []string
	versionManagementPluginPaths []string
	versionManagement map[string]version_management.VersionManagement
	continuousIntegration map[string]continuous_integration.ContinuousIntegration
)

type PluginReader struct {}

func init() {
	testingToolPluginPaths = getAllTestingToolPlugins()
	versionManagementPluginPaths = getAllVersionManagementFiles()
	versionManagement = make(map[string]version_management.VersionManagement)
	continuousIntegration = make(map[string]continuous_integration.ContinuousIntegration)
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func getAllTestingToolPlugins() (list []string) {
	if err := filepath.Walk("./plugins/continuous_integration", func(path string, info os.FileInfo, err error) error {
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

func getAllVersionManagementFiles() (list []string) {
	if err := filepath.Walk("./plugins/version_management", func(path string, info os.FileInfo, err error) error {
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

func (pr *PluginReader) LoadAllTestingToolPlugins() error {
	for _, p := range testingToolPluginPaths {
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
		pInterface, err := p.Lookup("NewTestingTools")
		if err != nil {
			log.Fatal(err)
			return err
		}
		newVersionManagement, _ := pInterface.(func() version_management.VersionManagement) // assert the type of the func
		versionManagement[derefString(pName.(*string))] = newVersionManagement()
	}
	return nil
}

func (pr *PluginReader) LoadAllVersionManagementPlugins() error {
	for _, p := range versionManagementPluginPaths {
		p, err := plugin.Open(p)
		if err != nil {
			log.Fatal(err)
			return err
		}
		pName, err := p.Lookup("PluginName")
		if err != nil {
			log.Fatal(err)
			return err
		}
		pInterface, err := p.Lookup("NewVersionManagement")
		if err != nil {
			log.Fatal(err)
			return err
		}
		newVersionManagement, _ := pInterface.(func() version_management.VersionManagement) // assert the type of the func
		versionManagement[derefString(pName.(*string))] = newVersionManagement()
	}
	return nil
}

func (pr *PluginReader) GetOrgaInfo(credentials map[string]interface{}) (info map[string]interface{}) {
	fmt.Println("Start GetOrgaInfo in PluginReader")
	info = make(map[string]interface{})
	for k, v := range versionManagement {
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