package main

import (
	"context"
	"fmt"
	heroku "github.com/heroku/heroku-go/v5"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"log"
)

var (
	PluginName string
	HerokuClient *heroku.Service
)

func init() {
	PluginName = "Heroku"
}

type Heroku struct {}

type Response struct{
	Addon struct {

	}
}

func (h *Heroku) SubmitCredentials(_, token string) {
	heroku.DefaultTransport.BearerToken = token
	HerokuClient = heroku.NewService(heroku.DefaultClient)
}

func (h *Heroku) FetchData() (string, error) {
	//var response Response
	addon, err := h.getAddonList()
	if err != nil {
		log.Println(err)
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

func (h *Heroku) getAddonList() (heroku.AddOnListResult, error){
	addon, err := HerokuClient.AddOnList(context.TODO(), &heroku.ListRange{Field: "name", Max: 1000})
	if err != nil {
		fmt.Println("\tin fetch error")
		return nil, err
	}
	return addon, nil
}

func (h *Heroku) get() {

}