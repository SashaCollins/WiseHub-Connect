/*
Interface for Datastore.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
 */
package data

type User struct {
	//gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Plugins  string `json:"plugins"`
}

type Plugin struct {
	//gorm.Model
	Name string `json:"name"`
	Token string `json:"token"`
	Description string `json:"description"`
}

type DatastoreI interface {
	Load(email ...string) (users User, plugin []Plugin, err error)
	Save(name string, password string, email string) error
	Update(option string, data ...string) error
	Delete(email string) error
}