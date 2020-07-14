/*
Interface for Listeners.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package listener

type Listener interface {
	GetOrgaInfo()
	GetTeamInfo()
	GetInsightTeamInfo()
	GetTeamRepoInfo()
}

