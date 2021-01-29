package viewmodel

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
)
var (
	pluginPaths []string
	pluginMap map[string]plugins.PluginI
)

type AdminView struct {
	Datastore data.DatastoreI
<<<<<<< HEAD
	PluginReader  plugins.PluginReader
	PluginI plugins.PluginI
}



func init() {
	pluginPaths = getAllPlugins()
	pluginMap = make(map[string]plugins.PluginI)

}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func getAllPlugins() (list []string) {
	if err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".so" {
			list = append(list, path)
		}
		return nil
	}); err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
=======
	PluginLoader  plugins.PluginLoader
>>>>>>> b2e13f6459b64687eef36c17e9654b8e26aef0d0
}

func (av *AdminView) LoadAllPlugins() error {
	for _, p := range pluginPaths {
		p, err := plugin.Open(p)
		fmt.Println(p)
		if err != nil {
			log.Fatal(err)
			return err
		}
		pName, err := p.Lookup("PluginName")
		if err != nil {
			log.Fatal(err)
			return err
		}
		pInterface, err := p.Lookup("NewPlugin")
		if err != nil {
			log.Fatal(err)
			return err
		}
		newPlugin, _ := pInterface.(func() plugins.PluginI) // assert the type of the func
		pluginMap[derefString(pName.(*string))] = newPlugin()
	}
	return nil
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
