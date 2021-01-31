package main

import (
	"context"
	"encoding/json"
	"github.com/drone/drone-go/drone"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"golang.org/x/oauth2"
	"log"
)

type Drone struct{}

var (
	PluginName   string
	droneClient  drone.Client
)


type Response struct {
	Repository struct {
		Name string `json:"repoName"`
		Branch string `json:"repoBranch"`
		Build struct {
			Number int64 `json:"buildNumber"`
			Status string `json:"buildStatus"`
		} `json:"build"`
	} `json:"repo"`
}

func init() {
	PluginName = "Drone CI"
}

func NewPlugin() plugins.PluginI {
	return &Drone{}
}

func (d *Drone) SubmitCredentials(host, token string) {
	oauthConfig := new(oauth2.Config)
	httpClient := oauthConfig.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: token,
		},
	)
	droneClient = drone.NewClient(host, httpClient)
}

func (d *Drone) FetchData() (string, error) {
	repos, err := droneClient.RepoList()
	if err != nil {
		log.Println("Data could not be fetched!")
		return "", err
	}
	var tmp Response
	var resp []Response
	for _, repo := range repos {
		tmp.Repository.Name = repo.Name
		tmp.Repository.Branch = repo.Branch
		build, _ := droneClient.BuildLast(repo.Namespace, repo.Name, repo.Branch)
		tmp.Repository.Build.Number = build.Number
		tmp.Repository.Build.Status = build.Status
		resp = append(resp, tmp)
	}
	r, _ := json.Marshal(resp)
	return string(r), err
}

func (d *Drone) FetchPluginName() string {
	return getPluginName()
}

func getPluginName() string {
	return PluginName
}