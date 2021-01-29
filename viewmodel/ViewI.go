/*
Interface for Views.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package viewmodel

import "github/SashaCollins/Wisehub-Connect/model/plugins"

type ViewI interface {
	GetData() ([]byte, error)
	SetPlugins(map[string]plugins.PluginI)
	SetCredentials(map[string]plugins.Credentials)
}