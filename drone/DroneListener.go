package drone

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strings"
	"github.com/julienschmidt/httprouter"
)

type DroneListener struct{
	dr droneReader
}

func (dl *DroneListener) GetOrgaInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := json.MarshalIndent(, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (dl *DroneListener) GetTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result, err := json.MarshalIndent(, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (dl *DroneListener) GetInsightTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result, err := json.MarshalIndent(, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (dl *DroneListener) GetTeamRepoInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result, err := json.MarshalIndent(*, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	secResult, err := json.MarshalIndent(*, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	result = append(result, secResult...)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (dl *DroneListener) StartServer(finished chan bool) {
	router := httprouter.New()
	router.GET("/courses", dl.GetOrgaInfo)
	router.GET("/courses/:orgaName", dl.GetTeamInfo)
	router.GET("/courses/:orgaName/:teamName", dl.GetInsightTeamInfo)
	router.GET("/courses/:orgaName/:teamName/:repoName", dl.GetInsightTeamInfo)
	log.Fatal(http.ListenAndServe(":8080", router))

	finished <- true
}
