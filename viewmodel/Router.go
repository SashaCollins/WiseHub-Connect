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
)

type Router struct {
	Datastore *data.Datastore
	View ViewI
}

var (
	PluginMap map[string]plugins.PluginI
)

type Response struct{
	Success 		bool                    	`json:"success"`
	Email 			string               		`json:"email"`
	Admin			bool 						`json:"admin"`
	Plugins 		[]data.Plugin           	`json:"plugins"`
	RepoData		map[string]interface{}		`json:"repos"`
	CourseData		map[string]interface{}		`json:"courses"`
}

type Request struct {
	Option 			string						`json:"option"`
	NewEmail 		string 						`json:"new_email"`
	Email 			string 						`json:"email"`
	Password 		string 						`json:"password"`
	Plugins 		[]data.Plugin 				`json:"plugins"`
	Repository 		string						`json:"repo"`
	Course 			string			 			`json:"course"`
}

func (r *Router) ReloadPlugins() map[string]plugins.PluginI {
	var loader PluginLoader
	pluginMap, err := loader.LoadAllPlugins()
	if err != nil {
		log.Fatal("Could not load plugins!")
	}
	return pluginMap
}

func (r *Router) LoadPluginCredentials(userEmail string) map[string]plugins.Credentials {
	credentialMap := make(map[string]plugins.Credentials)
	dbUser, err := r.Datastore.Load(userEmail)
	if err != nil {
		log.Fatal("User not found!")
	}
	for _, v := range dbUser[0].Plugins {
		credentialMap[v.PluginName] = plugins.Credentials{UserNameHost: v.UsernameHost, Token: v.Token}
	}
	return credentialMap
}

func (r *Router) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, _ := r.Datastore.Load(user.Email)
	if len(dbUser) == 0 {
		err = r.Datastore.Create(user.Password, user.Email)
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

func (r *Router) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := r.Datastore.Load(user.Email)
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

func (r *Router) Profile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var request Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Profile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		fmt.Printf("Profile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	dbUser, err := r.Datastore.Load(request.Email)
	if err != nil {
		fmt.Printf("Profile: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}

	response.Success = true
	response.Email = dbUser[0].Email
	response.Plugins = dbUser[0].Plugins
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Profile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (r *Router) Forgot(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := r.Datastore.Load(user.Email)
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

func (r *Router) Update(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var update Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Update: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = json.Unmarshal(reqBody, &update)
	if err != nil {
		fmt.Printf("Update: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	switch update.Option {
	case "email":
		dbUser, err := r.Datastore.Load(update.Email)
		if err != nil {
			fmt.Printf("Update: %v\n", dbUser)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["old"] = update.Email
		change["new"] = update.NewEmail
		if err = r.Datastore.Update(update.Option, change); err != nil {
			fmt.Printf("Update: %s\n", err)
			http.Error(w, "Invalid email", 668)
			return
		}
		response.Success = true
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("Update: %s\n", err)
			http.Error(w, "Internal server error", 500)
			return
		}
		_, _ = w.Write(resp)
		return
	case "password":
		if _, err := r.Datastore.Load(update.Email); err != nil {
			fmt.Printf("UpdatePassword: %v\n", err)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["email"] = update.Email
		change["password"] = update.Password
		if err = r.Datastore.Update(update.Option, change); err != nil {
			fmt.Printf("UpdatePassword: %s\n", err)
			http.Error(w, "Invalid email", 668)
			return
		}
		response.Success = true
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("UpdatePassword: %s\n", err)
			http.Error(w, "Internal server error", 500)
			return
		}
		_, _ = w.Write(resp)
		return
	case "credentials":
		change := make(map[string]interface{})
		change["email"] = update.Email
		change["updatedPlugins"] = update.Plugins
		if err := r.Datastore.Update(update.Option, change); err != nil {
			fmt.Printf("UpdatePlugins: %s\n", err)
			http.Error(w, "Invalid email", 668)
			return
		}
		response.Success = true
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("UpdatePlugins: %s\n", err)
			http.Error(w, "Internal server error", 500)
			return
		}
		_, _ = w.Write(resp)
		return
	default:
	}
}

func (r *Router) Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := r.Datastore.Load(request.Email)
	if err != nil {
		fmt.Printf("Show: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}

	credentials := r.LoadPluginCredentials(dbUser[0].Email)
	switch request.Option {
	case "general":
		r.View = &GeneralView{}
	case "template":
		//r.View = &TemplateView{}
		fmt.Println("Template")
	default:
		//r.View = &DefaultView{}
		fmt.Println("test")
	}
	r.View.SetPlugins(PluginMap)
	r.View.SetCredentials(credentials)
	fmt.Println(r.View)
	pluginData, err := r.View.GetData()
	if err != nil {
		fmt.Printf("Show: %v\n", dbUser)
		http.Error(w, "Invalid data", 668)
		return
	}
	fmt.Println(pluginData)

	//fmt.Println(PluginMap)
	//if len(dbUser) == 1 {
	//	for pName, pStruct := range PluginMap {
	//		fmt.Println(pName)
	//		fmt.Println(pStruct)
	//		//credentials["name"] = userPlugin.UsernameHost
	//		//credentials["token"] = userPlugin.Token
	//		//r.Plugins.SubmitCredentials(userPlugin.UsernameHost, userPlugin.Token)
	//		//courses, err = gv.Plugins.FetchData()}
	//	}
	//}

	//var view ViewI = NewGeneralView()

	response.Success = true

	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

//func (gv *GeneralView) Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	var user Request
//	var response Response
//
//	reqBody, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		fmt.Printf("DeleteProfile: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	if err := json.Unmarshal(reqBody, &user); err != nil {
//		fmt.Printf("DeleteProfile: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	if _, err := gv.Datastore.Load(user.Email); err != nil {
//		fmt.Printf("DeleteProfile: %v\n", err)
//		http.Error(w, "Invalid email", 668)
//		return
//	}
//	if err := gv.Datastore.Delete(user.Email); err != nil {
//		fmt.Printf("DeleteProfile: %s\n", err)
//		http.Error(w, "Invalid email", 668)
//		return
//	}
//	response.Success = true
//	resp, err := json.Marshal(response)
//	fmt.Println(resp)
//	if err != nil {
//		fmt.Printf("DeleteProfile: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	_, _ = w.Write(resp)
//	return
//}

func (r *Router) New() (router *httprouter.Router) {
	router = httprouter.New()

	// Authentication
	router.POST("/auth/signin", r.SignIn)
	router.POST("/auth/signup", r.SignUp)

	// Profile
	router.POST("/user/profile", r.Profile)
	router.POST("/user/update/password", r.Update)
	router.POST("/user/update/credentials", r.Update)

	// Fetch view data
	router.POST("/data/all", r.Show)
	//router.POST("/user/delete", r.View.Delete)
	//router.POST("/user/update/email", r.View.Update)
	//router.POST("/user/update/password", r.View.Update)
	//router.POST("/user/update/credentials", r.View.Update)
	//router.POST("/user/repos", r.View.Repositories)
	//router.POST("/user/all", r.view.Show())
	//router.POST("/user/teams", r.View.Teams)
	//router.POST("/admin/delete/plugins", r.View.Delete)
	//router.POST("/admin/update/plugins", r.View.Update)

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	return
}

func (r *Router) Run(port int, finished chan bool) {
	router := r.New()
	PluginMap = r.ReloadPlugins()
	fmt.Printf("Run: %s\n", http.ListenAndServe(fmt.Sprintf(":%d", port), router))
	finished <- true
}
