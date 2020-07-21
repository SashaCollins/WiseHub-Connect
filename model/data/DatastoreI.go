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
}

type DatastoreI interface {
	Load(args ...string) (users User ,err error)
	Save(name string, password string, email string) error
	Del(email string, password string ) error
}