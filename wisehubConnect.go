// credit - go-graphql hello world example
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github/SashaCollins/Wisehub-Connect/data"
	"log"
	//"github/SashaCollins/Wisehub-Connect/heroku"
	//gh "github/SashaCollins/Wisehub-Connect/github"
	//"github/SashaCollins/Wisehub-Connect/drone"
	//"github/SashaCollins/Wisehub-Connect/heroku"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	fmt.Println("start")

	fmt.Println("1")
	ds := data.Datastore{}
	fmt.Println("1")
	ds.Start()
	fmt.Println("1")

	//githubFinished := make(chan bool)
	//gl := gh.GithubListener{}
	//go gl.StartServer(githubFinished)
	//fmt.Println("\tgithub running...")
	//droneFinished := make(chan bool)
	//dl := drone.DroneListener{}
	//go dl.StartServer(droneFinished)
	//fmt.Println("\tdrone running...")
	//herokuFinished := make(chan bool)
	//hl := heroku.HerokuListener{}
	//go hl.StartServer(herokuFinished)
	//fmt.Println("\theroku running...")
	//
	//<- githubFinished
	//<- droneFinished
	//<- herokuFinished

	fmt.Println("end")
}
