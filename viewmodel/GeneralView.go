/*
This is the general view which collects all relevant data from all plugins
@author SashaCollins
@version 1.0
 */
package viewmodel

import (
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"log"
)

type GeneralView struct {
	Plugin map[string]plugins.PluginI
	Credentials map[string]plugins.Credentials
}

func (gv *GeneralView) GetData() (response map[string]string, err error) {
	response = make(map[string]string)
	for pName, pValue := range gv.Credentials {
		extension := gv.Plugin[pName]
		if extension == nil {
			continue
		}
		extension.SubmitCredentials(pValue.UserNameHost, pValue.Token)
		var data string
		data, err = extension.FetchData()
		if err != nil {
			log.Println("Data could not be fetched!")
			continue
		}
		response[extension.FetchPluginName()] = data
	}
	return response, err
}

func (gv *GeneralView) SetPlugins(plugin map[string]plugins.PluginI) {
	gv.Plugin = plugin
}

func (gv *GeneralView) SetCredentials(credentials map[string]plugins.Credentials) {
	gv.Credentials = credentials
}
