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

type NormalView struct {
	Datastore data.DatastoreI
	Reader  plugins.Reader
}

func (nv *NormalView) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, _ := nv.Datastore.Load(user.Email)
	if len(dbUser) == 0 {
		err = nv.Datastore.Save(user.Password, user.Email)
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

func (nv *NormalView) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("SignIn: %v\n", dbUser)
		http.Error(w, "Invalid email or password", 667)
		return
	}
	if dbUser[0].Password == user.Password {
		response.Success = true
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

func (nv *NormalView) Forgot(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := nv.Datastore.Load(user.Email)
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

func (nv *NormalView) Update(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
		dbUser, err := nv.Datastore.Load(update.Email)
		if err != nil {
			fmt.Printf("Update: %v\n", dbUser)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["old"] = update.Email
		change["new"] = update.NewEmail
		if err = nv.Datastore.Update(update.Option, change); err != nil {
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
		if _, err := nv.Datastore.Load(update.Email); err != nil {
			fmt.Printf("UpdatePassword: %v\n", err)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["email"] = update.Email
		change["password"] = update.Password
		if err = nv.Datastore.Update(update.Option, change); err != nil {
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
	case "plugins":
		change := make(map[string]interface{})
		change["email"] = update.Email
		change["updatedPlugins"] = update.Plugins
		if err := nv.Datastore.Update(update.Option, change); err != nil {
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

func (nv *NormalView) Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	if err := json.Unmarshal(reqBody, &user); err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	if _, err := nv.Datastore.Load(user.Email); err != nil {
		fmt.Printf("DeleteProfile: %v\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}
	if err := nv.Datastore.Delete(user.Email); err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}
	response.Success = true
	resp, err := json.Marshal(response)
	fmt.Println(resp)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, err := nv.Datastore.Load(request.Email)
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

func (nv *NormalView) Repositories(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
		fmt.Printf("Repositories: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	dbUser, err := nv.Datastore.Load(request.Email)
	if err != nil {
		fmt.Printf("Repositories: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}
	fmt.Println(dbUser)

	organizations := nv.Reader.GetOrgaInfo()
	jsonString, err := json.Marshal(organizations)
	if err != nil {
		fmt.Printf("Repositories: %v\n", jsonString)
		http.Error(w, "Internal server error", 500)
		return
	}
	response.Organization = jsonString
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

func (nv *NormalView) Courses(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user Request
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}









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

func (nv *NormalView) Run(port int, finished chan bool) {
	router := Router{View: nv}
	normalRouter := router.New()
	fmt.Printf("Run: %s\n", http.ListenAndServe(fmt.Sprintf(":%d", port), normalRouter))
	finished <- true
}