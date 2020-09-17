package main

import (
	"fmt"
	"github.com/bgentry/heroku-go"
	"github/SashaCollins/Wisehub-Connect/model/config"
)

var herokuClient heroku.Client

func init(){
	conf := config.GetConfig()
	email := conf.Heroku.Username
	apiToken := conf.Heroku.APIToken
	herokuClient = heroku.Client{Username: email, Password: apiToken}
}

type HerokuReader struct {}

func (hr *HerokuReader) FetchData(info int) (interface{}, error){
	switch info {
	case 1:
		fmt.Println("\tin fetch")
		var apps []heroku.App
		// pass nil for options if you don't need to set any optional params
		app, err := herokuClient.AppCreate(nil)
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
			// if this is a heroku.Error, it will contain details about the error
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
			// if this is a heroku.Error, it will contain details about the error
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

var Heroku HerokuReader
