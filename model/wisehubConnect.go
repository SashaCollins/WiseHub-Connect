// credit - go-graphql hello world example
package main

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/viewmodel"
)

//func init() {
//	// loads values from .env into the system
//	if err := godotenv.Load(); err != nil {
//		log.Print("No .env file found")
//	}
//}
func main() {
	fmt.Println("start")
	//ds := data.Datastore{}
	//err := ds.Save("name092","pw092","email092")
	//if err != nil {
	//	fmt.Printf("\tmain: %s\n", err)
	//}
	//_, _, err = ds.Load("email092")
	//if err != nil {
	//	fmt.Printf("\tmain: %s\n", err)
	//}

	//githubFinished := make(chan bool)
	//gl := gh.GithubListener{}
	//go gl.StartServer(githubFinished)
	//fmt.Println("\tgithub running...")
	//droneFinished := make(chan bool)
	//dl := testing_tools.DroneListener{}
	//go dl.StartServer(droneFinished)
	//fmt.Println("\tdrone running...")
	//herokuFinished := make(chan bool)
	//hl := deployment.HerokuListener{}
	//go hl.StartServer(herokuFinished)
	//fmt.Println("\theroku running...")
	//
	//<- githubFinished
	//<- droneFinished
	//<- herokuFinished

	ds := data.Datastore{}
	routerFinished := make(chan bool)
	router := viewmodel.Router{Datastore: &ds}
	router.Run(9010, routerFinished)
	<- routerFinished



	//personalViewFinished := make(chan bool)
	//av := viewmodel.PersonalView{Datastore: &ds}
	//if err := av.LoadAllPlugins(); err != nil {
	//	panic("Could not load Plugins")
	//}
	//go av.Run(9020, personalViewFinished)
	//<- personalViewFinished

	//generalViewFinished := make(chan bool)
	//nv := viewmodel.GeneralView{Datastore: &ds}
	//go nv.Run(9010, generalViewFinished)
	//<- generalViewFinished

	fmt.Println("end")
}
