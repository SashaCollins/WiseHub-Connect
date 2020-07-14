// credit - go-graphql hello world example
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	//"github/SashaCollins/Wisehub-Connect/model/data"
	"log"
	//"github/SashaCollins/Wisehub-Connect/heroku"
	//gh "github/SashaCollins/Wisehub-Connect/github"
	//"github/SashaCollins/Wisehub-Connect/drone"
	//"github/SashaCollins/Wisehub-Connect/heroku"
	"github/SashaCollins/Wisehub-Connect/viewmodel"

)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	fmt.Println("start")
	//ds := data.Datastore{}
	//ds.Save()
	//ds.Load("name092","pw092","email092")

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

	normalViewFinished := make(chan bool)
	nv := viewmodel.NormalView{}
	go nv.Run(normalViewFinished)
	<- normalViewFinished


	fmt.Println("end")
}
