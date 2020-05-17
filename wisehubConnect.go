// credit - go-graphql hello world example
package main

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/config"
	"log"
	gh "github/SashaCollins/Wisehub-Connect/github"
	"sync"

	//"github/SashaCollins/Wisehub-Connect/drone"
	//"github/SashaCollins/Wisehub-Connect/heroku"
	"github.com/joho/godotenv"
)

var lock sync.Mutex
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	fmt.Println("start")
	//lock.Lock()
	conf := config.GetConfig()
	//lock.Unlock()
	fmt.Println(conf.GitHub.Username)
	fmt.Println( conf.GitHub.APIToken)
	//fmt.Println(conf.DebugMode)
	//fmt.Println(conf.MaxUsers)

	// Print out each role
	//for _, role := range conf.UserRoles {
	//	fmt.Println(role)
	//}

	githubFinished := make(chan bool)
	gl := gh.GithubListener{}
	go gl.StartServer(githubFinished)
	<- githubFinished
	//droneFinished := make(chan bool)
	//dl := drone.DroneListener{}
	//go dl.StartServer(droneFinished)
	//<- droneFinished
	//herokuFinished := make(chan bool)
	//hl := heroku.HerokuListener{}
	//go hl.StartServer(herokuFinished)
	//<- herokuFinished

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
	fmt.Println("end")
}
