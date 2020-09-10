/*
Interface for Datastore.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
 */
package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     		string
	Password 		string
	Email    		string		`gorm:"not null;unique"`
	Plugins  		[]Plugin	`gorm:"foreignKey:UserID"`
}

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
	Load(email ...string) (user User, err error)
	Save(name, password, email string) error
	Update(option string, data map[string]interface{}) error
	Delete(email string) error
}