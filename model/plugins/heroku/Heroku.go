package main

import (
	"context"
	"fmt"
	heroku "github.com/heroku/heroku-go/v5"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
)

var (
	PluginName string
	HerokuClient *heroku.Service
)

func init() {
	PluginName = "Heroku"
}

type Heroku struct {}

func (h *Heroku) SubmitCredentials(_, token string) {
	heroku.DefaultTransport.BearerToken = token
	HerokuClient = heroku.NewService(heroku.DefaultClient)
}

func (h *Heroku) FetchData() (string, error) {
	account, err := HerokuClient.AccountInfo(context.TODO())
	if err != nil {
		fmt.Println("\tin fetch error")
		return "", err
	}
	fmt.Print(account)
	addon, err := HerokuClient.AddOnList(context.TODO(), nil)
	if err != nil {
		fmt.Println("\tin fetch error")
		return "", err
	}
	fmt.Println(addon)
	return "", nil
}

func (h *Heroku) FetchPluginName() string {
	return getPluginName()
}

func getPluginName() string {
	return PluginName
}

func init(){
	PluginName = "Heroku"
}

func NewPlugin() plugins.PluginI {
	return &Heroku{}
}