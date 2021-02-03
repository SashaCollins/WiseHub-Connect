/*
Interface for Plugins.
This is important to ensure modularity.
Do not edit or delete.
Must be implemented by every plugin in order to work.
For more information please read ./template/Template.go
@author SashaCollins
*/
package plugins
/*
To consistently convey credentials
 */
type Credentials struct {
	UserNameHost string
	Token string
}

type PluginI interface {
	/*
	The Plugin uses these credentials to fetch the data from the respective API
	 */
	SubmitCredentials(username string, token string)
	/*
	Returns all Data which will be displayed in the dashboard.
	To ensure readability the return string should match a (custom) struct, which is convertible to JSON.
	 */
	FetchData() (string, error)
	/*
	Returns the name of the Plugin for visualization purposes
	 */
	FetchPluginName() string
}