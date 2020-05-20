package drone

import (
	"context"
	"fmt"
	"github.com/drone/drone-go/drone"
	"github/SashaCollins/Wisehub-Connect/config"
	"golang.org/x/oauth2"
)

type droneReader struct{}
var (
	client drone.Client
	currentBuild build
	currentRepo repo
)
func init() {
	// create an http client with oauth authentication
	conf := config.GetConfig()
	Host := conf.DroneCI.Host
	DroneToken := conf.DroneCI.APIToken
	oauthConfig := new(oauth2.Config)
	httpClient := oauthConfig.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: DroneToken,
		},
	)
	// create the drone client with authenticator
	client = drone.NewClient(Host, httpClient)
	// Use client...
}
type step struct {
	Name string
	Status string
}
type stage struct {
	Name string
	Kind string
	Status string
	Steps []step
}
type build struct {
	//Name string
	//Owner string
	Branch string
	Status string
	Time int64
	Stages []stage
}
type repo struct {
	Name string
	Owner string
	Branch string
	Build build
}
func (dr *droneReader) jsonBuildToStructBuild(target *build, source drone.Build) *build {
	target.Branch = source.Ref
	target.Status = source.Status
	target.Time = source.Started
	if len(source.Stages) > 0 {
		for _, sourceStage := range source.Stages {
			var targetStage stage
			targetStage.Name = sourceStage.Name
			targetStage.Kind = sourceStage.Kind
			targetStage.Status = sourceStage.Status

			for _, sourceStep := range sourceStage.Steps {
				var targetStep step
				targetStep.Name = sourceStep.Name
				targetStep.Status = sourceStep.Status
				targetStage.Steps = append(targetStage.Steps, targetStep)
			}
			target.Stages = append(target.Stages, targetStage)
		}
	}
	return target
}
func (dr *droneReader) fetchData(info int) (interface{}, error){
	switch info {
	case 1:
		// gets the current user
		user, err := client.Self()
		return user, err

	case 2:
		// gets the named repository information
		repo, err := client.Repo("WiseHub-Connector", "WiseHub-Project")
		if err != nil {
			return nil, fmt.Errorf("something went wrong with collecting the data")
		}
		currentRepo.Name = repo.Name
		currentRepo.Owner = repo.Namespace
		currentRepo.Branch = repo.Branch
		currentRepo.Build = *dr.jsonBuildToStructBuild(&currentRepo.Build, repo.Build)
		return currentRepo, nil

	case 3:
		buildLast, err := client.BuildLast("WiseHub-Connector", "WiseHub-Project", "master")
		if err != nil {
			return nil, fmt.Errorf("something went wrong with collecting the data")
		}
		currentBuild = *dr.jsonBuildToStructBuild(&currentBuild, *buildLast)
		return currentBuild, nil
	default:
		return nil, fmt.Errorf("something went wrong with the info number %s", info)
	}
}

