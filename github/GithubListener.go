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

func (gl *GithubListener) GetOrgaInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var viewer = *gl.gr.getViewer()
	var allOrgas = *gl.gr.getOrganizations(viewer.Viewer.Login)
	result, err := json.MarshalIndent(allOrgas, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (gl *GithubListener) GetTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var orgaName = ps.ByName("orgaName")
	var allTeams = *gl.gr.getTeamsPerOrganization((githubv4.String)(orgaName))
	result, err := json.MarshalIndent(allTeams, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (gl *GithubListener) GetInsightTeamInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var orgaName = ps.ByName("orgaName")
	var teamName = ps.ByName("teamName")
	var allTeamMembersAndRepos = *gl.gr.getTeamMembersAndRepositories((githubv4.String)(orgaName), (githubv4.String)(teamName))
	result, err := json.MarshalIndent(allTeamMembersAndRepos, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (gl *GithubListener) GetTeamRepoInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var viewer = *gl.gr.getViewer()
	var repoName = ps.ByName("repoName")
	var repoOwner = "" //TODO
	var allIssuesAssigned, allCommits = gl.gr.getRepositoryInfo((githubv4.String)(repoName), (githubv4.String)(repoOwner), viewer.Viewer.Login)
	result, err := json.MarshalIndent(*allIssuesAssigned, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	secResult, err := json.MarshalIndent(*allCommits, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	result = append(result, secResult...)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (gl *GithubListener) StartServer(finished chan bool) {
	router := httprouter.New()
	router.GET("/courses", gl.GetOrgaInfo)
	router.GET("/courses/:orgaName", gl.GetTeamInfo)
	router.GET("/courses/:orgaName/:teamName", gl.GetInsightTeamInfo)
	router.GET("/courses/:orgaName/:teamName/:repoName", gl.GetInsightTeamInfo)
	log.Fatal(http.ListenAndServe(":8080", router))

	finished <- true
}
