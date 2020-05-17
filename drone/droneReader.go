package drone

import (
	"context"
	"fmt"
	"github/SashaCollins/Wisehub-Connect/config"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

type droneReader struct{}
var client drone.Client
func init() {
	// create an http client with oauth authentication.
	conf := config.New()
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
func (gr *droneReader) fetchData(info int) (interface{}, error){
	switch info {

	case 1:
		// gets the current user
		user, err := client.Self()
		fmt.Println(user, err)
		return user, err

	case 2:
		// gets the named repository information
		repo, err := client.Repo("drone", "drone-go")
		fmt.Println(repo, err)
		return repo, err

	default:
		return nil, fmt.Errorf("something went wrong with the info number %s", info)
	}
}

