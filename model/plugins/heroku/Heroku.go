package main

import (
	"fmt"
	heroku "github.com/heroku/heroku-go/v5"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
)

var (
	usaername string
	password string
	PluginName string
	herokuClient *heroku.Service
)

type Heroku struct {}

func init(){
	PluginName = "Heroku"
}

func NewTestingTools() plugins.PluginI {
	return &Heroku{}
}

func (h *Heroku) UpdateCredentials(credentials map[string]string) {
	heroku.DefaultTransport.Username = credentials["username"]
	heroku.DefaultTransport.Password = credentials["token"]
	herokuClient = heroku.NewService(heroku.DefaultClient)
}

func (h *Heroku) fetchData(info int) (interface{}, error){
	switch info {
	case 1:
		fmt.Println("\tin fetch")
		var apps []heroku.App
		// pass nil for options if you don't need to set any optional params
		app, err := herokuClient.InvoiceInfo()
		if err != nil {
			fmt.Println("\tin fetch error")
			return nil, err
		}
		fmt.Println("Created", app.Name)
		apps = append(apps, *app)
		return apps, err
		// Output:
		// Created dodging-samurai-42

	case 2:
		var apps []heroku.App
		name := "myapp"
		region := "region"

		// Optional values need to be provided as pointers. If a field in an option
		// struct is nil (not provided), the option is omitted from the API request.
		opts := heroku.AppCreateOpts{Name: &name, Region: &region}

		// Create an app with options set:
		app2, err := herokuClient.AppCreate(&opts)
		if err != nil {
			// if this is a deployment.Error, it will contain details about the error
			if hkerr, ok := err.(heroku.Error); ok {
				return nil, fmt.Errorf("Error id=%s message=%q", hkerr.Id, hkerr)
			}
			return nil, err
		}
		fmt.Printf("created app2: name=%s region=%s", app2.Name, app2.Region.Name)
		apps = append(apps, *app2)
		return apps, err

		// Output:
		// created app2: name=myapp region=eu

	case 3:
		apps, err := herokuClient.AppList(&heroku.ListRange{Field: "name", Max: 1000})
		if err != nil {
			// if this is a deployment.Error, it will contain details about the error
			if hkerr, ok := err.(heroku.Error); ok {
				return nil, fmt.Errorf("Error id=%s message=%q", hkerr.Id, hkerr)
			}
			return nil, err
		}
		fmt.Println(apps)
		return apps, err

	default:
		return nil, fmt.Errorf("something went wrong with the info number %s", info)
	}
}

func (h *Heroku) GetRepositories() (interface{}, error) {

	return nil, nil
}

func (h *Heroku) GetBuilds() (interface{}, error) {

	return nil, nil
}

func (h *Heroku) GetOrgaInfo(i interface{}) (interface{}, error) {
	panic("implement me")
}

func (h *Heroku) GetTeamInfo(orgaName string) (interface{}, error) {
	panic("implement me")
}

func (h *Heroku) GetInsightTeamInfo(orgaName, teamName string) (interface{}, error) {
	panic("implement me")
}

func (h *Heroku) GetTeamRepoInfo(repoName, repoOwner string) (interface{}, interface{}, error) {
	panic("implement me")
}