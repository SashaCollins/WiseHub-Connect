// credit - go-graphql hello world example
package main

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/heroku"
	"log"
	gh "github/SashaCollins/Wisehub-Connect/github"
	"github/SashaCollins/Wisehub-Connect/drone"
	//"github/SashaCollins/Wisehub-Connect/heroku"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	fmt.Println("start")
	//fmt.Println(conf.DebugMode)
	//fmt.Println(conf.MaxUsers)

	// Print out each role
	//for _, role := range conf.UserRoles {
	//	fmt.Println(role)
	//}

	githubFinished := make(chan bool)
	gl := gh.GithubListener{}
	go gl.StartServer(githubFinished)
	fmt.Println("\tgithub running...")
	droneFinished := make(chan bool)
	dl := drone.DroneListener{}
	go dl.StartServer(droneFinished)
	fmt.Println("\tdrone running...")
	herokuFinished := make(chan bool)
	hl := heroku.HerokuListener{}
	go hl.StartServer(herokuFinished)
	fmt.Println("\theroku running...")

	<- githubFinished
	<- droneFinished
	<- herokuFinished

	fmt.Println("end")
}
