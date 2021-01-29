package viewmodel

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"io/ioutil"
	"net/http"
)

type AdminView struct {
	Datastore data.DatastoreI
	PluginLoader  plugins.PluginLoader
}
func (av *AdminView) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("SignIn: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("SignIn: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	dbUser, err := av.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("SignIn: %v\n", dbUser)
		http.Error(w, "Invalid email or password", 667)
		return
	}
	if dbUser[0].Password == user.Password {
		response.Success = true
		response.Admin = dbUser[0].Admin
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("SignIn: %s\n", err)
			http.Error(w, "Internal server error", 500)
		}
		_, _ = w.Write(resp)
	} else {
		fmt.Printf("SignIn: %s\n", err)
		http.Error(w, "Invalid email or password", 667)
	}
	return
}
func (av *AdminView) Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var request Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	dbUser, err := av.Datastore.Load(request.Email)
	if err != nil {
		fmt.Printf("Show: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}


	response.Success = true
	response.Email = dbUser[0].Email
	response.Plugins = dbUser[0].Plugins
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (av *AdminView) Run(port int, finished chan bool) {
	router := Router{View: av}
	adminRouter := router.New()
	fmt.Printf("Run: %s\n", http.ListenAndServe(fmt.Sprintf(":%d", port), adminRouter))
	finished <- true
}
