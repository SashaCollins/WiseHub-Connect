package viewmodel

import (
	"encoding/json"
	//"encoding/json"
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

//type User struct{
//	Name string `json:"name"`
//	Password string `json:"password"`
//	Email string `json:"email"`
//}
type NormalView struct {
	Datastore data.DatastoreI
	Listener  listener.Listener
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
		fmt.Printf("ERROR: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	//fmt.Printf("%s\n", user)
	//fmt.Println(nv.Datastore)
	err = nv.Datastore.Save(user.Name, user.Password, user.Email)
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response.Success = true
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
		_, _ = w.Write(resp)
	} else {
		fmt.Printf("ERROR: %s\n", err)
		http.Error(w, "User already exists", 666)
	}
	return
}
func (nv *NormalView) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user data.User
	var response Response

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	dbUser, err := nv.Datastore.Load(user.Email)
	if err == nil && dbUser.Password == user.Password {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response.Success = true
		resp, err := json.Marshal(response)
		fmt.Println(resp)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
		_, _ = w.Write(resp)
	} else {
		fmt.Printf("ERROR: %s\n", err)
		http.Error(w, "Invalid email or password", 667)
	}
	return
}
func (nv *NormalView) Show() {

}
func (nv *NormalView) Run(finished chan bool) {
	router := httprouter.New()
	router.POST("/auth/signup", nv.SignUp)
	router.POST("/auth/signin", nv.SignIn)

	fmt.Printf("ERROR: %s\n", http.ListenAndServe(":9010", router))
	finished <- true
}