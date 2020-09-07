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

type User struct {
	//gorm.Model
	ID 		 		uint		`gorm:"primaryKey;not null"`
	Name     		string
	Password 		string
	Email    		string		`gorm:"unique;not null"`
	Plugins  		[]Plugin
}

type Plugin struct {
	//gorm.Model
	ID 				uint		`gorm:"primaryKey;not null"`
	UserNameHost 	string
	Token 			string
	Description 	string		`gorm:"unique;not null"`
	Updated 		bool
}

type DatastoreI interface {
	New(driver string, data map[string]interface{}) (db *gorm.DB)
	Load(email ...string) (users User)
	Save(name, password, email string) error
	Update(option string, data ...interface{}) error
	Delete(email string) error
}