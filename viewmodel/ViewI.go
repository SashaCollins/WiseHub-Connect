/*
Interface for Views.
This is important to ensure modularity.
do not edit or delete.
Must be implemented by every new View.
For more information please read ./TemplateView.go
@author SashaCollins
@version 1.0
*/
package viewmodel

import "github/SashaCollins/Wisehub-Connect/model/plugins"

type ViewI interface {
	/*
	Fetch all Data from all Plugins
	 */
	GetData() (map[string]string, error)
	/*
	Setter for plugins
	used for plugin access
	 */
	SetPlugins(map[string]plugins.PluginI)
	/*
	Setter for the credentials for the plugins
	to convey them
	 */
	SetCredentials(map[string]plugins.Credentials)
}