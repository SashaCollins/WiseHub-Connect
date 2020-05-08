package main

import (

)
type Listener interface {
	GetOrgaInfo()
	GetTeamInfo()
	GetInsightTeamInfo()
	GetTeamRepoInfo()
}

