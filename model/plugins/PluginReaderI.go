/*
Interface for Listeners.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package plugins

type PluginReaderI interface {
	// Plugins
	LoadAllTestingToolPlugins() error
	LoadAllVersionManagementPlugins() error

	// Version Management
	GetOrgaInfo() (info map[string]interface{})
	GetTeamInfo() (info map[string]interface{})
	GetInsightTeamInfo() (info map[string]interface{})
	GetTeamRepoInfo() (info map[string]interface{})

	// Testing Tools
	GetRepositories() (info map[string]interface{})
	GetBuilds() (info map[string]interface{})

}