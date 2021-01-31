/*
Interface for Listeners.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package plugins

type Credentials struct {
	UserNameHost string
	Token string
}

type PluginI interface {
	SubmitCredentials(username, token string)
	FetchData() (string, error)
	FetchPluginName() string
}