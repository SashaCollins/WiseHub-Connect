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
	Reader  plugins.Reader
}

func (av *AdminView) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	dbUser, _ := av.Datastore.Load(user.Email)
	if len(dbUser) == 0 {
		err = av.Datastore.Save(user.Password, user.Email)
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

func (av *AdminView) Forgot(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Forgot: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("Forgot: %s\n", err)
	}

	dbUser, err := av.Datastore.Load(user.Email)
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

func (av *AdminView) Update(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
		dbUser, err := av.Datastore.Load(update.Email)
		if err != nil {
			fmt.Printf("Update: %v\n", dbUser)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["old"] = update.Email
		change["new"] = update.NewEmail
		if err = av.Datastore.Update(update.Option, change); err != nil {
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
		if _, err := av.Datastore.Load(update.Email); err != nil {
			fmt.Printf("UpdatePassword: %v\n", err)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["email"] = update.Email
		change["password"] = update.Password
		if err = av.Datastore.Update(update.Option, change); err != nil {
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
		if err := av.Datastore.Update(update.Option, change); err != nil {
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

func (av *AdminView) Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
	if _, err := av.Datastore.Load(user.Email); err != nil {
		fmt.Printf("DeleteProfile: %v\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}
	if err := av.Datastore.Delete(user.Email); err != nil {
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

func (av *AdminView) Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
	dbUser, err := av.Datastore.Load(user.Email)
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

func (av *AdminView) TestAdmin(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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
	dbUser, err := av.Datastore.Load(user.Email)
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

func (av *AdminView) Repositories(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	repos := av.Reader.GetOrgaInfo()

	fmt.Println(repos)








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

func (av *AdminView) Courses(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

func (av *AdminView) Run(port int, finished chan bool) {
	router := Router{View: av}
	adminRouter := router.New()
	adminRouter.POST("/admin/test", av.TestAdmin)
	fmt.Printf("Run: %s\n", http.ListenAndServe(fmt.Sprintf(":%d", port), adminRouter))
	finished <- true
}