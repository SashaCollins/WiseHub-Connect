package logging

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (gl *GithubListener) StartServer(finished chan bool) {
	router := httprouter.New()
	router.GET("/settings", gl.GetOrgaInfo)
	router.GET("/settings/logging:orgaName", gl.GetTeamInfo)
	log.Fatal(http.ListenAndServe(":8080", router))

	finished <- true
}
