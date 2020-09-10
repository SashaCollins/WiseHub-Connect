package viewmodel

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/model/listener"
	"io/ioutil"
	"net/http"
)

type NormalView struct {
	Datastore data.DatastoreI
	Listener  listener.Listener
}

func (nv *NormalView) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user User
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
	if dbUser.Email != "" || dbUser.Email == user.Email {
		fmt.Printf("SignUp: %s\n", err)
		http.Error(w, "User already exists!", 666)
		return
	}

	err = nv.Datastore.Save(user.Name, user.Password, user.Email)
	if err != nil {
		fmt.Println("3: User already exists!")
		http.Error(w, "User already exists!", 666)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func (nv *NormalView) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user User
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
	if dbUser.Password == user.Password {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response.Success = true
		resp, err := json.Marshal(response)
		fmt.Println(resp)
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

	dbUser, err := nv.Datastore.Load(user.Email)
	fmt.Println(dbUser)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func (nv *NormalView) UpdateEmail(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var email UpdateEmail
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = json.Unmarshal(reqBody, &email)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	dbUser, err := nv.Datastore.Load(email.OldEmail)
	if err != nil {
		fmt.Printf("UpdateEmail: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}

	update := make(map[string]interface{})
	update["old"] = email.OldEmail
	update["new"] = email.NewEmail
	err = nv.Datastore.Update("email", update)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

func (nv *NormalView) UpdatePassword(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user User
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	if _, err := nv.Datastore.Load(user.Email); err != nil {
		fmt.Printf("UpdatePassword: %v\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}

	update := make(map[string]interface{})
	update["email"] = user.Email
	update["password"] = user.Password
	if err = nv.Datastore.Update("password", update); err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) UpdatePlugins(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var plugins UpdatePlugins
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("UpdatePlugins: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = json.Unmarshal(reqBody, &plugins)
	if err != nil {
		fmt.Printf("UpdatePlugins: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	update := make(map[string]interface{})
	update["email"] = plugins.Email
	update["plugins"] = plugins.Plugins
	if err := nv.Datastore.Update("plugins", update); err != nil {
		fmt.Printf("UpdatePlugins: %s\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("UpdatePlugins: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) DeleteProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user User
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	var user data.User
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

	dbUser, err := nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("Show: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	response.Email = dbUser.Email
	response.Plugins = dbUser.Plugins
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) Update(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
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

	dbUser, err := nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("Show: %v\n", dbUser)
		http.Error(w, "Invalid email", 668)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	response.Email = dbUser.Email
	response.Plugins = dbUser.Plugins
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) Run(finished chan bool) {
	router := httprouter.New()
	router.POST("/auth/signup", nv.SignUp)
	router.POST("/auth/signin", nv.SignIn)
	router.POST("/user/profile", nv.Show)
	router.POST("/user/update/email", nv.UpdateEmail)
	router.POST("/user/update/password", nv.UpdatePassword)
	router.POST("/user/delete", nv.DeleteProfile)
	router.POST("/user/update/plugins", nv.UpdatePlugins)

	fmt.Printf("Run: %s\n", http.ListenAndServe(":9010", router))
	finished <- true
}