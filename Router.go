package WiseHub_Connect

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
type Router struct {
	Router *httprouter.Router
	Handler map[string]httprouter.Handle
}
func (nv *Router) init() {
	//ds := data.Datastore{}
	//nv.Datastore = &ds
}

func (nv *Router) Run(finished chan bool) {
	router := httprouter.New()
	//router.POST("/auth/signup", nv.SignUp)
	//router.POST("/auth/signin", nv.SignIn)
	//router.POST("/auth/signin", nv.SignIn)
	//router.POST("/auth/signin", nv.SignIn)
	//router.POST("/auth/signin", nv.SignIn)
	//router.POST("/auth/signin", nv.SignIn)
	//router.POST("/auth/signin", nv.SignIn)

	fmt.Printf("ERROR: %s\n", http.ListenAndServe(":9010", router))
	finished <- true
}