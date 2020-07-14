package viewmodel

import (
	//"encoding/json"
	"fmt"
	"github/SashaCollins/Wisehub-Connect/model/listener"
	_ "github/SashaCollins/Wisehub-Connect/model/listener"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct{
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type NormalView struct {
	ds data.DatastoreI
	l  listener.Listener
}
func (nv *NormalView) signUp(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	_, _ = fmt.Fprintf(w, "%v: %v\n", req.Body, req.Form)
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reqBody)
	fmt.Printf("%v\n", reqBody)

	//fmt.Println(req.Body)
	//decoder := json.NewDecoder(req.Body)
	//fmt.Println(decoder)
	//var user User
	//err = decoder.Decode(&user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(user)

	//var user tmp
	//err = json.Unmarshal(reqBody, &user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", user)
	//nv.ds.Save()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, _ = w.Write(reqBody)
}
func (nv *NormalView) signIn(w http.ResponseWriter, req *http.Request) {

}
func (nv *NormalView) Show() {

}
func (nv *NormalView) Run(finished chan bool) {
	http.HandleFunc("/auth/signup", nv.signUp)
	//http.HandleFunc("/signIn", nv.signIn)

	http.ListenAndServe(":9010", nil)
	finished <- true
}