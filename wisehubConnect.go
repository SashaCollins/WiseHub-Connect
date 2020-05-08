// credit - go-graphql hello world example
package main

import (
	"fmt"
	gh "wisehubConnect/github"
)


func show(item interface{}) {
	gh.PrintJSON(item)
}

func main() {
	finished := make(chan bool)
	gl := gh.GithubListener{}
	go gl.StartServer(finished)
	<- finished

	//fmt.Println("before run")
	//var viewer = *gh.GetViewer()
	//show(viewer)
	////printJSON(currentViewer)
	//var allOrgas = *gh.GetOrganizations(viewer.Viewer.Login)
	//show(allOrgas)
	//var allTeams = *gh.GetTeamsPerOrganization(allOrgas[0].Login)
	//show(allTeams)
	//var allTeamMembersAndRepos = *gh.GetTeamMembersAndRepositories(allOrgas[0].Login, allTeams[0].Slug)
	//show(allTeamMembersAndRepos)
	//var allIssuesAssigned, allRefs = gh.GetRepositoryInfo(allTeamMembersAndRepos.Organization.Team.Repositories.Nodes[0].Name, allTeamMembersAndRepos.Organization.Team.Repositories.Nodes[0].Owner.Login, viewer.Viewer.Login)
	//show(allIssuesAssigned)
	//show(allRefs)

	fmt.Println("#############################################")
	//printJSON(currentUser)
	//printJSON(allOrganizations)
	fmt.Println("after run")
}
