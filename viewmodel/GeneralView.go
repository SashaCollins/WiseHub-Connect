package viewmodel

import (
	"fmt"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"log"
)

type GeneralView struct {
	Plugin map[string]plugins.PluginI
	Credentials map[string]plugins.Credentials
}

//func NewGeneralView () *GeneralView {
//	gv := &AbstractView{}
//	ggv := &GeneralView{plugin, gv}
//	gv.ViewI = ggv
//	return ggv
//}

func (gv *GeneralView) GetData() (map[string]string, error) {
	response := make(map[string]string)
	for pName, pValue := range gv.Credentials {
		extension := gv.Plugin[pName]
		if extension == nil {
			continue
		}
		extension.SubmitCredentials(pValue.UserNameHost, pValue.Token)
		data, err := extension.FetchData()
		if err != nil {
			log.Fatal("Data could not be fetched!")
			return nil, err
		}
		fmt.Println(data)
		response[extension.FetchPluginName()] = data
	}
	return response, nil
}

func (gv *GeneralView) SetPlugins(plugin map[string]plugins.PluginI) {
	gv.Plugin = plugin
}

func (gv *GeneralView) SetCredentials(credentials map[string]plugins.Credentials) {
	gv.Credentials = credentials
}



//func (gv *GeneralView) Courses(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	var request Request
//	var response Response
//
//	reqBody, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	err = json.Unmarshal(reqBody, &request)
//	if err != nil {
//		fmt.Printf("Repositories: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	//only 1 user possible
//	dbUser, err := gv.Datastore.Load(request.Email)
//	if err != nil {
//		fmt.Printf("Repositories: %v\n", dbUser)
//		http.Error(w, "Invalid email", 668)
//		return
//	}
//	credentials := make(map[string]string)
//	courses := make(map[string]interface{})
//
//	if len(dbUser) == 1 {
//		for _, userPlugin := range dbUser[0].Plugins {
//			credentials["name"] = userPlugin.UsernameHost
//			credentials["token"] = userPlugin.Token
//			gv.Plugins.SubmitCredentials(userPlugin.UsernameHost, userPlugin.Token)
//			courses, err = gv.Plugins.FetchData()
//		}
//	}
//
//	//courses = nv.PluginLoader.GetOrgaInfo(user.PluginName, credentials)
//
//	response.CourseData = courses
//	fmt.Println("End Courses in NormalView")
//	response.Success = true
//	resp, err := json.Marshal(response)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	_, _ = w.Write(resp)
//	return
//}
//
//func (gv *GeneralView) Repositories(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	var user Request
//	var response Response
//
//	reqBody, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	err = json.Unmarshal(reqBody, &user)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//
//
//
//
//
//
//
//
//
//	response.Success = true
//	resp, err := json.Marshal(response)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	_, _ = w.Write(resp)
//	return
//}
//
//func (gv *GeneralView) Teams(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	var user Request
//	var response Response
//
//	reqBody, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	err = json.Unmarshal(reqBody, &user)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//
//
//
//
//
//
//
//
//
//	response.Success = true
//	resp, err := json.Marshal(response)
//	if err != nil {
//		fmt.Printf("Show: %s\n", err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	_, _ = w.Write(resp)
//	return
//}