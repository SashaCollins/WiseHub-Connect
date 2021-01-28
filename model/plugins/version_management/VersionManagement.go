package version_management

type VersionManagement interface {
	GetOrgaInfo(interface{}) (interface{}, error)
	GetTeamInfo(orgaName string) (interface{}, error)
	GetInsightTeamInfo(orgaName, teamName string) (interface{}, error)
	GetTeamRepoInfo(repoName, repoOwner string) (interface{}, interface{}, error)
}
