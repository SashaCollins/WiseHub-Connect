/*
Interface for Listeners.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package plugins

type Credentials struct {
	UserNameHost string
	Token string
}

type PluginI interface {
	SubmitCredentials(username, token string)
	FetchData() (string, error)
	FetchPluginName() string

	// Version Management
	//GetOrgaInfo(interface{}) (interface{}, error)
	//GetTeamInfo(string) (interface{}, error)
	//GetInsightTeamInfo(string, string) (interface{}, error)
	//GetTeamRepoInfo(string, string) (interface{}, interface{}, error)
	//GetRepositoryInfo() (interface{}, error)
	// Testing Tools
	//GetRepositoryInfo() (interface{}, error)
	//GetBuilds() (interface{}, error)

}