package drone

import (
	"context"
	"fmt"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)
const (
	token = ""
	host  = "http://localhost:9080"
)
var client drone.Client
func init() {
	// create an http client with oauth authentication.

	config := new(oauth2.Config)
	httpClient := config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: token,
		},
	)
	// create the drone client with authenticator
	client = drone.NewClient(host, httpClient)
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

