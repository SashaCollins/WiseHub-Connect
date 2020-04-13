// credit - go-graphql hello world example
package main

import (
	gh "wisehubConnect/github"
	"fmt"
)

func show(item interface{}) {
	gh.PrintJSON(item)
}

func main() {
	fmt.Println("before run")
	var viewer = *gh.GetViewer()
	show(viewer)
	//printJSON(currentViewer)
	var allTeams = *gh.GetTeams()
	show(allTeams)


	fmt.Println("#############################################")
	//printJSON(currentUser)
	//printJSON(allOrganizations)
	fmt.Println("after run")
}
