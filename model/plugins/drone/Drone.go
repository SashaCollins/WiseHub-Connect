package main

import (
	"context"
	"fmt"
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
		Name string `json:"repo_name"`
		Branch string `json:"repo_branch"`
		Build struct {
			Number int64 `json:"build_number"`
			Status string `json:"build_status"`
		} `json:"build"`
	} `json:"repo"`
}

func init() {
	PluginName = "Drone CI"
}

func NewPlugin() plugins.PluginI {
	return &Drone{}
}

func (d *Drone) FetchSomething() error {
	panic("implement me")
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

func (d *Drone) FetchData() (interface{}, error) {
	fmt.Println("in Drone CI FetchData")
	var tmp Response
	var response []Response
	repos, err := droneClient.RepoList()
	if err != nil {
		log.Fatal("Data could not be fetched!")
	}
	for _, repo := range repos {
		tmp.Repository.Name = repo.Name
		tmp.Repository.Branch = repo.Branch
		build, _ := droneClient.BuildLast(repo.Namespace, repo.Name, repo.Branch)
		tmp.Repository.Build.Number = build.Number
		tmp.Repository.Build.Status = build.Status
		response = append(response, tmp)
	}
	fmt.Println(response)
	return response, err
}

func (d *Drone) FetchPluginName() string {
	return getPluginName()
}

func getPluginName() string {
	return PluginName
}

//func (d *Drone) jsonBuildToStructBuild(target *Build, source drone.Build) *Build {
//	target.Branch = source.Ref
//	target.Status = source.Status
//	target.Time = source.Started
//	if len(source.Stages) > 0 {
//		for _, sourceStage := range source.Stages {
//			var targetStage Stage
//			targetStage.Name = sourceStage.Name
//			targetStage.Kind = sourceStage.Kind
//			targetStage.Status = sourceStage.Status
//
//			for _, sourceStep := range sourceStage.Steps {
//				var targetStep Step
//				targetStep.Name = sourceStep.Name
//				targetStep.Status = sourceStep.Status
//				targetStage.Steps = append(targetStage.Steps, targetStep)
//			}
//			target.Stages = append(target.Stages, targetStage)
//		}
//	}
//	return target
//}

//func (d *Drone) fetchData(info int) (interface{}, error) {
//	switch info {
//	case 1:
//		// gets the current user
//		user, err := droneClient.Self()
//		return user, err
//
//	case 2:
//		// gets the named repository information
//		repo, err := droneClient.Repo("WiseHub-Connector", "WiseHub-Project")
//		if err != nil {
//			return nil, fmt.Errorf("something went wrong with collecting the data")
//		}
//		currentRepo.Name = repo.Name
//		currentRepo.Owner = repo.Namespace
//		currentRepo.Branch = repo.Branch
//		currentRepo.Build = *d.jsonBuildToStructBuild(&currentRepo.Build, repo.Build)
//		return currentRepo, nil
//
//	case 3:
//		buildLast, err := droneClient.BuildLast("WiseHub-Connector", "WiseHub-Project", "master")
//		if err != nil {
//			return nil, fmt.Errorf("something went wrong with collecting the data")
//		}
//		currentBuild = *d.jsonBuildToStructBuild(&currentBuild, *buildLast)
//		return currentBuild, nil
//	default:
//		return nil, fmt.Errorf("something went wrong with the info number %s", info)
//	}
//}