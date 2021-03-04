/*
This is a template to how a concrete view should look like, which collects all relevant data from all plugins.
A new View e.g. PersonalView should somehow differentiate from general view.
You should either trim or extend received data from FetchData Method.
Or the plugins should extend a new method (e.g. FetchSpecialData, FetchExtendedData, FetchShortenedData) for different data
@author SashaCollins
@version 1.0
*/
package viewmodel

import (
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"log"
)

type TemplateView struct {
	Plugin map[string]plugins.PluginI
	Credentials map[string]plugins.Credentials
}

func (tv *TemplateView) GetData() (map[string]string, error) {
	response := make(map[string]string)
	for pName, pValue := range tv.Credentials {
		extension := tv.Plugin[pName]
		if extension == nil {
			continue
		}
		extension.SubmitCredentials(pValue.UserNameHost, pValue.Token)
		data, err := extension.FetchData()
		if err != nil {
			log.Println("Data could not be fetched!")
			continue
		}
		response[extension.FetchPluginName()] = data
	}
	return response, nil
}

func (tv *TemplateView) SetPlugins(plugin map[string]plugins.PluginI) {
	tv.Plugin = plugin
}

func (tv *TemplateView) SetCredentials(credentials map[string]plugins.Credentials) {
	tv.Credentials = credentials
}
