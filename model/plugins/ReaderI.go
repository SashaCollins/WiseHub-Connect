/*
Interface for Listeners.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package plugins

type ReaderI interface {
	GetOrgaInfo() map[string]interface{}
	GetTeamInfo(orgaName string) map[string]interface{}
	GetInsightTeamInfo(orgaName, teamName string) map[string]interface{}
	GetTeamRepoInfo(repoName, repoOwner string) map[string]interface{}
}