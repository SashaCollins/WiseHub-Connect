package heroku

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	_ "strings"
)

type HerokuListener struct{
	hr herokuReader
}

func (hl *HerokuListener) GetOrgaInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := hl.hr.fetchData(1)
	result, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (hl *HerokuListener) GetTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//result, err := json.MarshalIndent(, "", "\t")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	http.Error(w, "internal error", http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(result)
}
func (hl *HerokuListener) GetInsightTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//result, err := json.MarshalIndent(, "", "\t")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	http.Error(w, "internal error", http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(result)
}
func (hl *HerokuListener) GetTeamRepoInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//result, err := json.MarshalIndent(*, "", "\t")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	http.Error(w, "internal error", http.StatusInternalServerError)
	//	return
	//}
	//secResult, err := json.MarshalIndent(*, "", "\t")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	http.Error(w, "internal error", http.StatusInternalServerError)
	//	return
	//}
	//result = append(result, secResult...)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(result)
}

func (hl *HerokuListener) StartServer(finished chan bool) {
	router := httprouter.New()
	router.GET("/courses", hl.GetOrgaInfo)
	router.GET("/courses/:orgaName", hl.GetTeamInfo)
	router.GET("/courses/:orgaName/:teamName", hl.GetInsightTeamInfo)
	router.GET("/courses/:orgaName/:teamName/:repoName", hl.GetInsightTeamInfo)
	log.Fatal(http.ListenAndServe(":6080", router))

	finished <- true
}
