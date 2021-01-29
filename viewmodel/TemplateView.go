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

//var (
//	pluginPaths []string
//	pluginMap map[string]plugins.PluginI
//)

type PersonalView struct {
	Datastore data.DatastoreI
	PluginI plugins.PluginI
}



//func init() {
//	pluginPaths = getAllPlugins()
//	pluginMap = make(map[string]plugins.PluginI)
//
//}

//func derefString(s *string) string {
//	if s != nil {
//		return *s
//	}
//	return ""
//}

//func getAllPlugins() (list []string) {
//	if err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
//		if info.IsDir() {
//			return nil
//		}
//		if filepath.Ext(path) == ".so" {
//			list = append(list, path)
//		}
//		return nil
//	}); err != nil {
//		fmt.Printf("walk error [%v]\n", err)
//	}
//	return list
//
//}

//func (pv *PersonalView) LoadAllPlugins() error {
//	for _, p := range pluginPaths {
//		p, err := plugin.Open(p)
//		fmt.Println(p)
//		if err != nil {
//			log.Fatal(err)
//			return err
//		}
//		pName, err := p.Lookup("PluginName")
//		if err != nil {
//			log.Fatal(err)
//			return err
//		}
//		pInterface, err := p.Lookup("NewPlugin")
//		if err != nil {
//			log.Fatal(err)
//			return err
//		}
//		newPlugin, _ := pInterface.(func() plugins.PluginI) // assert the type of the func
//		pluginMap[derefString(pName.(*string))] = newPlugin()
//	}
//	return nil
//}



func (pv *PersonalView) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("SignUp: %s\n", err)
		http.Error(w, "Internal server error", 500)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("SignUp: %s\n", err)
		http.Error(w, "Internal server error", 500)
	}

	dbUser, _ := pv.Datastore.Load(user.Email)
	if len(dbUser) == 0 {
		err = pv.Datastore.Create(user.Password, user.Email)
		if err != nil {
			fmt.Println("3: User already exists!")
			http.Error(w, "User already exists!", 666)
			return
		}

		response.Success = true
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("SignUp: %s\n", err)
			http.Error(w, "Internal server error", 500)
			return
		}
		_, _ = w.Write(resp)
		return
	}
	if dbUser[0].Email != "" || dbUser[0].Email == user.Email {
		fmt.Printf("SignUp: %s\n", err)
		http.Error(w, "User already exists!", 666)
		return
	}
}

func (pv *PersonalView) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := pv.Datastore.Load(user.Email)
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

func (pv *PersonalView) Forgot(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Forgot: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("Forgot: %s\n", err)
	}

	dbUser, err := pv.Datastore.Load(user.Email)
	fmt.Println(dbUser)
	response.Success = true
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (pv *PersonalView) Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := pv.Datastore.Load(request.Email)
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

//func (pv *PersonalView) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
//
//}
//
//func (pv *PersonalView) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
//
//}

//func (pv *PersonalView) Run(port int, finished chan bool) {
//	router := Router{View: pv}
//	personalRouter := router.New()
//	//personalRouter.POST("/admin/delete/plugins", pv.Delete)
//	//personalRouter.POST("/admin/update/plugins", pv.Update)
//	fmt.Printf("Run: %s\n", http.ListenAndServe(fmt.Sprintf(":%d", port), personalRouter))
//	finished <- true
//}


