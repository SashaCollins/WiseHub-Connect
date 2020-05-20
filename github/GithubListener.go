package github

import (
	"encoding/json"
	"fmt"
	"github.com/shurcooL/githubv4"
	"log"
	"net/http"
	_ "strings"
	"github.com/julienschmidt/httprouter"
)

type GithubListener struct{
	gr githubReader
}

func (gl *GithubListener) GetHello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := json.MarshalIndent("Hello there!", "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (gl *GithubListener) GetOrgaInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	viewer, err := gl.gr.getViewer()
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	allOrgas, err := gl.gr.getOrganizations(viewer.Viewer.Login)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	result, err := json.MarshalIndent(allOrgas, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (gl *GithubListener) GetTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	orgaName := ps.ByName("orgaName")
	allTeams, err := gl.gr.getTeamsPerOrganization((githubv4.String)(orgaName))
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	result, err := json.MarshalIndent(allTeams, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (gl *GithubListener) GetInsightTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	orgaName := ps.ByName("orgaName")
	teamName := ps.ByName("teamName")
	allTeamMembersAndRepos, err := gl.gr.getTeamMembersAndRepositories((githubv4.String)(orgaName), (githubv4.String)(teamName))
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	result, err := json.MarshalIndent(allTeamMembersAndRepos, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (gl *GithubListener) GetTeamRepoInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	viewer, err := gl.gr.getViewer()
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	repoName := ps.ByName("repoName")
	repoOwner := "" //TODO
	allIssuesAssigned, allCommits, err := gl.gr.getRepositoryInfo((githubv4.String)(repoName), (githubv4.String)(repoOwner), viewer.Viewer.Login)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	result, err := json.MarshalIndent(*allIssuesAssigned, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	secResult, err := json.MarshalIndent(*allCommits, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	result = append(result, secResult...)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (gl *GithubListener) StartServer(finished chan bool) {
	router := httprouter.New()
	router.GET("/", gl.GetHello)
	router.GET("/courses", gl.GetOrgaInfo)
	router.GET("/courses/:orgaName", gl.GetTeamInfo)
	router.GET("/courses/:orgaName/:teamName", gl.GetInsightTeamInfo)
	router.GET("/courses/:orgaName/:teamName/:repoName", gl.GetTeamRepoInfo)
	log.Fatal(http.ListenAndServe(":8080", router))

	finished <- true
}
