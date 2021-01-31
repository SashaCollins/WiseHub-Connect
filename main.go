package main

import (
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/viewmodel"
)

func main() {
	ds := data.Datastore{}
	routerFinished := make(chan bool)
	router := viewmodel.Router{Datastore: &ds}
	router.Run(9010, routerFinished)
	<- routerFinished
}
