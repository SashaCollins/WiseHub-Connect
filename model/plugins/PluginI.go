/*
Interface for Listeners.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package plugins

type PluginI interface {
	SubmitCredentials(username string, token string)

	// Version Management
	GetOrgaInfo(interface{}) (interface{}, error)
	GetTeamInfo(string) (interface{}, error)
	GetInsightTeamInfo(string, string) (interface{}, error)
	GetTeamRepoInfo(string, string) (interface{}, interface{}, error)

	// Testing Tools
	GetRepositories() (interface{}, error)
	GetBuilds() (interface{}, error)

}