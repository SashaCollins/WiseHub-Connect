package drone

import (
	"context"
	"fmt"
	"github/SashaCollins/Wisehub-Connect/config"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

var client drone.Client
func init() {
	// create an http client with oauth authentication.
	conf := config.New()
	DroneToken := conf.DroneCI.APIToken
	Host := conf.DroneCI.Host
	config := new(oauth2.Config)
	httpClient := config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: DroneToken,
		},
	)
	// create the drone client with authenticator
	client = drone.NewClient(Host, httpClient)
	// Use client...
}
type droneReader struct{}

func fetchData() {
	// gets the current user
	user, err := client.Self()
	fmt.Println(user, err)

	// gets the named repository information
	repo, err := client.Repo("drone", "drone-go")
	fmt.Println(repo, err)
}

