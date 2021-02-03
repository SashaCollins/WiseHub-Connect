/*
Interface for Datastore.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
 */
package data

import (
	"gorm.io/gorm"
)
/*
This struct is used to create the User table in the database
 */
type User struct {
	gorm.Model
	Email    		string		`gorm:"not null;unique"`
	Password 		string
	Plugins  		[]Plugin	`gorm:"foreignKey:UserID"`
}
/*
This struct is used to create the Plugin table in the database
 */
type Plugin struct {
	gorm.Model
	UserID       	uint
	PluginName		string
	UsernameHost 	string
	Token 			string
	Description 	string
	Updated 		bool
}

type DatastoreI interface {
	/*
	Load an user from the database
	Input either zero or one argument
	zero loads all users from database
	one loads user based on email address
	 */
	Load(email ...string) (user []User, err error)
	/*
	Save a new user in the database
	 */
	Create(password, email string) error
	/*
	Update any user information
	@param option - which information is going to be updated, e.g. email, password, credentials, ...
	@param data - the new information
	 */
	Update(option string, data map[string]interface{}) error
	/*
	Delete an user from database
	 */
	Delete(email string) error

}