package main

import (
	"context"
	"fmt"
	heroku "github.com/heroku/heroku-go/v5"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
)

var (
	PluginName string
	herokuClient *heroku.Service
)
type Response struct {
	Repository struct {
		Name string `json:"repo_name"`
		URL string
		Team string `json:"team_name"`
		Build struct {
			Status string `json:"build_status"`
		} `json:"build"`
	} `json:"repo"`
}

type Heroku struct {}

func init() {
	PluginName = "Heroku"
}

func NewPlugin() plugins.PluginI {
	return &Heroku{}
}
func getPluginName() string {
	return PluginName
}

func (h *Heroku) FetchData() (string, error) {
	account, err := herokuClient.AccountInfo(context.Background())
	if err != nil {
		fmt.Println("\tin fetch error")
		return "", err
	}
	fmt.Println("Heroku Info")
	fmt.Println(*account)
	appList, err := herokuClient.AppList(context.Background(), nil)
	if err != nil {
		fmt.Println("\tin fetch error")
		return "", err
	}
	for _, app := range appList {
		fmt.Println(app.Name)
		buildList, err := herokuClient.BuildList(context.Background(), app.ID, nil)
		if err != nil {
			fmt.Println("\tin fetch error")
			return "", err
		}
		fmt.Println(app.WebURL)
		fmt.Println(app.GitURL)
		fmt.Println(app.Owner)
		for _, build := range buildList {
			fmt.Println(build.Status)
			fmt.Println(build.UpdatedAt)
			fmt.Println(build.CreatedAt)
		}

	}

	return "", nil
}

func (h *Heroku) SubmitCredentials(email string, token string) {
	heroku.DefaultTransport.BearerToken = token
	//heroku.DefaultTransport.Username = email
	//heroku.DefaultTransport.Password = token
	herokuClient = heroku.NewService(heroku.DefaultClient)
}
func (h *Heroku) FetchPluginName() string {
	return getPluginName()
}


