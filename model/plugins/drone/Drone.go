package main

import (
	"context"
	"fmt"
	"github.com/drone/drone-go/drone"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"golang.org/x/oauth2"
)

type Drone struct{}

var (
	PluginName   string
	droneClient  drone.Client
	currentBuild Build
	currentRepo  Repo
)

func init() {
	PluginName = "Drone"
	// create an http client with oauth authentication
	Host := conf.DroneCI.Host
	DroneToken := conf.DroneCI.APIToken
	oauthConfig := new(oauth2.Config)
	httpClient := oauthConfig.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: DroneToken,
		},
	)
	// create the testing_tools client with authenticator
	droneClient = drone.NewClient(Host, httpClient)
	// Use client...
}

func NewPlugin() plugins.PluginI {
	return &Drone{}
}

type Step struct {
	Name string
	Status string
}

type Stage struct {
	Name string
	Kind string
	Status string
	Steps []Step
}

type Build struct {
	//Name string
	//Owner string
	Branch string
	Status string
	Time int64
	Stages []Stage
}

type Repo struct {
	Name   string
	Owner  string
	Branch string
	Build  Build
}

func (d *Drone) jsonBuildToStructBuild(target *Build, source drone.Build) *Build {
	target.Branch = source.Ref
	target.Status = source.Status
	target.Time = source.Started
	if len(source.Stages) > 0 {
		for _, sourceStage := range source.Stages {
			var targetStage Stage
			targetStage.Name = sourceStage.Name
			targetStage.Kind = sourceStage.Kind
			targetStage.Status = sourceStage.Status

			for _, sourceStep := range sourceStage.Steps {
				var targetStep Step
				targetStep.Name = sourceStep.Name
				targetStep.Status = sourceStep.Status
				targetStage.Steps = append(targetStage.Steps, targetStep)
			}
			target.Stages = append(target.Stages, targetStage)
		}
	}
	return target
}

func (d *Drone) fetchData(info int) (interface{}, error) {
	switch info {
	case 1:
		// gets the current user
		user, err := droneClient.Self()
		return user, err

	case 2:
		// gets the named repository information
		repo, err := droneClient.Repo("WiseHub-Connector", "WiseHub-Project")
		if err != nil {
			return nil, fmt.Errorf("something went wrong with collecting the data")
		}
		currentRepo.Name = repo.Name
		currentRepo.Owner = repo.Namespace
		currentRepo.Branch = repo.Branch
		currentRepo.Build = *d.jsonBuildToStructBuild(&currentRepo.Build, repo.Build)
		return currentRepo, nil

	case 3:
		buildLast, err := droneClient.BuildLast("WiseHub-Connector", "WiseHub-Project", "master")
		if err != nil {
			return nil, fmt.Errorf("something went wrong with collecting the data")
		}
		currentBuild = *d.jsonBuildToStructBuild(&currentBuild, *buildLast)
		return currentBuild, nil
	default:
		return nil, fmt.Errorf("something went wrong with the info number %s", info)
	}
}
func (d *Drone) NewPlugin() plugins.PluginI {
	panic("implement me")
}

func (d *Drone) GetOrgaInfo(that interface{}) (interface{}, error) {
	panic("implement me")
}

func (d *Drone) GetTeamInfo(string) (interface{}, error) {
	panic("implement me")
}

func (d *Drone) GetInsightTeamInfo(string, string) (interface{}, error) {
	panic("implement me")
}

func (d *Drone) GetTeamRepoInfo(string, string) (interface{}, interface{}, error) {
	panic("implement me")
}
func (d *Drone) GetRepositories() (interface{}, error) {

	return nil, nil
}

func (d *Drone) GetBuilds() (interface{}, error) {

	return nil, nil
}