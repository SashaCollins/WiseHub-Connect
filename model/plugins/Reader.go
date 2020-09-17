package plugins

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

var (
	plugins []string
)

type Reader struct {}

func init() {
	plugins = getAllFiles()
}

func getAllFiles() (list []string) {
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

func (r *Reader) GetOrgaInfo() map[string]interface{} {
	mapOrgaInfo := make(map[string]interface{})
	for index := range plugins {
		p := plugins[index]
		fmt.Println(p)
		plug, err := plugin.Open(p)
		if err != nil {
			continue
		}
		symbol, err := plug.Lookup("GetOrgaInfo")
		if err != nil {
			fmt.Printf("GetOrgaInfo: Lookup error {%v}\n", err)
			continue
		}
		// symbol – Checks the function signature
		getOrgaInfo, ok := symbol.(func() (string, interface{}))
		if !ok {
			panic("GetOrgaInfo: Plugin has no 'GetOrgaInfo()map[string]interface{}' function")
		}
		fmt.Println(getOrgaInfo())
		pluginName, organization := getOrgaInfo()
		mapOrgaInfo[pluginName] = organization
	}
	return mapOrgaInfo
}

func (r *Reader) GetTeamInfo(orgaName string) map[string]interface{} {
	mapTeamInfo := make(map[string]interface{})
	for index := range plugins {
		p := plugins[index]
		plug, err := plugin.Open(p)
		if err != nil {
			continue
		}
		symbol, err := plug.Lookup("GetTeamInfo")
		if err != nil {
			fmt.Println("GetTeamInfo: Lookup error")
			continue
		}
		// symbol – Checks the function signature
		getTeamInfo, ok := symbol.(func(orgaName string) (string, interface{}))
		if !ok {
			fmt.Println("GetTeamInfo: Plugin has no 'GetOrgaInfo()map[string]interface{}' function")
			continue
		}
		fmt.Println(getTeamInfo(orgaName))
		pluginName, shortTeam := getTeamInfo(orgaName)
		mapTeamInfo[pluginName] = shortTeam
	}
	return mapTeamInfo
}

func (r *Reader) GetInsightTeamInfo(orgaName, teamName string) map[string]interface{} {
	mapInsightTeamInfo := make(map[string]interface{})
	for index := range plugins {
		p := plugins[index]
		plug, err := plugin.Open(p)
		if err != nil {
			continue
		}
		symbol, err := plug.Lookup("GetInsightTeamInfo")
		if err != nil {
			fmt.Println("GetInsightTeamInfo: Lookup error")
			continue
		}
		// symbol – Checks the function signature
		getInsightTeamInfo, ok := symbol.(func(orgaName, teamName string) (string, interface{}))
		if !ok {
			fmt.Println("GetInsightTeamInfo: Plugin has no 'GetOrgaInfo()map[string]interface{}' function")
			continue
		}
		fmt.Println(getInsightTeamInfo(orgaName, teamName))
		pluginName, team := getInsightTeamInfo(orgaName, teamName)
		mapInsightTeamInfo[pluginName] = team
	}
	return mapInsightTeamInfo
}

func (r *Reader) GetTeamRepoInfo(repoName, repoOwner string) map[string]interface{} {
	mapTeamRepoInfo := make(map[string]interface{})
	mapValue := make(map[string]interface{})
	for index := range plugins {
		p := plugins[index]
		plug, err := plugin.Open(p)
		if err != nil {
			continue
		}
		symbol, err := plug.Lookup("GetTeamRepoInfo")
		if err != nil {
			fmt.Println("GetTeamInfo: Lookup error")
			continue
		}
		// symbol – Checks the function signature
		getTeamRepoInfo, ok := symbol.(func(repoName, repoOwner string) (string, interface{}, interface{}))
		if !ok {
			fmt.Println("GetTeamInfo: Plugin has no 'GetOrgaInfo()map[string]interface{}' function")
			continue
		}
		fmt.Println(getTeamRepoInfo(repoName, repoOwner))
		pluginName, issue, ref := getTeamRepoInfo(repoName, repoOwner)
		mapValue["issue"] = issue
		mapValue["ref"] = ref
		mapTeamRepoInfo[pluginName] = mapValue
	}
	return mapTeamRepoInfo
}