package viewmodel

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/model/listener"
	_ "github/SashaCollins/Wisehub-Connect/model/listener"
	"io/ioutil"
	"net/http"
)

var (
	//ds data.Datastore
)

type NormalView struct {
	Datastore data.DatastoreI
	Listener  listener.Listener
}

type UpdateEmail struct {
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

func (nv *NormalView) init() {
	//ds := data.Datastore{}
	//nv.Datastore = &ds
}

func (nv *NormalView) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
	var response Response

	//fmt.Println(req.Body)
	//_, _ = fmt.Fprintf(w, "%v: %v\n", req.Body, req.Form)
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("SignUp: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("SignUp: %s\n", err)
	}
	_, _, err = nv.Datastore.Load(user.Email)
	//fmt.Printf("%s\n", user)
	//fmt.Println(nv.Datastore)
	err = nv.Datastore.Save(user.Name, user.Password, user.Email)
	if err != nil {
		fmt.Printf("SignUp: %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("SignUp: %s\n", err)
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("SignIn: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("SignIn: %s\n", err)
	}

	dbUser, _, err := nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("SignIn: %s\n", err)
	}
	if err == nil && dbUser.Password == user.Password {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response.Success = true
		resp, err := json.Marshal(response)
		fmt.Println(resp)
		if err != nil {
			fmt.Printf("SignIn: %s\n", err)
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

	dbUser, _, err := nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("Forgot: %s\n", err)
	}
	if err == nil && dbUser.Password == user.Password {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response.Success = true
		resp, err := json.Marshal(response)
		fmt.Println(resp)
		if err != nil {
			fmt.Printf("Forgot: %s\n", err)
		}
		_, _ = w.Write(resp)
	} else {
		fmt.Printf("Forgot: %s\n", err)
		http.Error(w, "Invalid email or password", 667)
	}
	return
}

func (nv *NormalView) UpdateEmail(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var email UpdateEmail
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &email)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
	}

	_, _, err = nv.Datastore.Load(email.OldEmail)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
	}
	err = nv.Datastore.Update("email", email.NewEmail, email.OldEmail)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	resp, err := json.Marshal(response)
	fmt.Println(resp)
	if err != nil {
		fmt.Printf("UpdateEmail: %s\n", err)
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) UpdatePassword(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
	}

	_, _, err = nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
		http.Error(w, "Invalid email", 668)
	}
	err = nv.Datastore.Update("password", user.Password, user.Email)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
	}


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	resp, err := json.Marshal(response)
	fmt.Println(resp)
	if err != nil {
		fmt.Printf("UpdatePassword: %s\n", err)
	}
	_, _ = w.Write(resp)
	return
}

func (nv *NormalView) DeleteProfile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
	}

	_, _, err = nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Invalid email", 668)
	}

	err = nv.Datastore.Delete(user.Email)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	resp, err := json.Marshal(response)
	fmt.Println(resp)
	if err != nil {
		fmt.Printf("DeleteProfile: %s\n", err)
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
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
	}

	dbUser, dbPlugins, err := nv.Datastore.Load(user.Email)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
		http.Error(w, "Invalid email", 668)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response.Success = true
	response.Email = dbUser.Email
	response.Plugins = dbPlugins
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Show: %s\n", err)
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

	fmt.Printf("Run: %s\n", http.ListenAndServe(":9010", router))
	finished <- true
}