package heroku

import (
	"fmt"
	"github.com/bgentry/heroku-go"
)
func init(){
	client := heroku.Client{Username: "email@me.com", Password: "my-api-key"}

	// pass nil for options if you don't need to set any optional params
	app, err := client.AppCreate(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created", app.Name)

	// Output:
	// Created dodging-samurai-42

	name := "myapp"
	region := "region"

	// Optional values need to be provided as pointers. If a field in an option
	// struct is nil (not provided), the option is omitted from the API request.
	opts := heroku.AppCreateOpts{Name: &name, Region: &region}

	// Create an app with options set:
	app2, err := client.AppCreate(&opts)
	if err != nil {
		// if this is a heroku.Error, it will contain details about the error
		if hkerr, ok := err.(heroku.Error); ok {
			panic(fmt.Sprintf("Error id=%s message=%q", hkerr.Id, hkerr))
		}
	}
	fmt.Printf("created app2: name=%s region=%s", app2.Name, app2.Region.Name)

	// Output:
	// created app2: name=myapp region=eu
	apps, err = client.AppList(&heroku.ListRange{Field: "name", Max: 1000})
	fmt.Println(apps)
}
type herokuReader struct {}

func fetchData() {
	
}
