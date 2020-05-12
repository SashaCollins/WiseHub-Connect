// credit - go-graphql hello world example
package main

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/config"
	"log"
	gh "github/SashaCollins/Wisehub-Connect/github"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	// Print out environment variables
	fmt.Println(conf.GitHub.Username)
	fmt.Println(conf.GitHub.APIToken)
	fmt.Println(conf.DebugMode)
	//fmt.Println(conf.MaxUsers)

	// Print out each role
	//for _, role := range conf.UserRoles {
	//	fmt.Println(role)
	//}

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
