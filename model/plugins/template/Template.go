/*
This is a template to show how to write a custom Plugin.
Do not edit or delete.
Every plugin must have its own folder within the plugins folder.
Please note that in order for shared objects to work,
every plugin.go file must be in main package.
For shared objects to work please include these line in start_wisehub.sh
go build -buildmode=plugin -o model/plugins/template/template.so model/plugins/template/Template.go
PLEASE CHANGE 'TEMPLATE' TO THE NAME OF YOUR PLUGIN
@author SashaCollins
@version 1.0
 */

package main

import (
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"go/doc"
)

var (
	pluginName string
	templateClient doc.Example.Client
)
/*
Your response should look something like this
It should contain every information you want displayed in the dashboard
 */
type Response struct {
	Repository struct {
		Name string `json:"repoName"`
		URL string
		Team string `json:"teamName"`
		Build struct {
			Status string `json:"buildStatus"`
		} `json:"build"`
	} `json:"repo"`
}

type Template struct {}

func init() {
	pluginName = "Template"
}
/*
Every Plugin must implement this method
for shared objects to create a .so file
 */
func NewPlugin() plugins.PluginI {
	return &Template{}
}
func getPluginName() string {
	return pluginName
}
/*
Ensure for readability purposes that the return data matches a (custom) response struct, which is convertible to JSON.
 */
func (h *Template) FetchData() (string, error) {
	var response string

	/*
	TODO
	fetch some data from templateClient
	 */

	return response, nil
}
/*
This Method should start the client connection to the respective API with the provided credentials.
 */
func (h *Template) SubmitCredentials(username string, token string) {
	/*
	TODO
	start client connection
	 */
}

func (h *Template) FetchPluginName() string {
	return getPluginName()
}

