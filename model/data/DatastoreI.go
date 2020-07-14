/*
Interface for Datastore.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
 */
package data

import "database/sql"

type User struct {
	Id int
	Name string
	Password string
	Email string
}

type DatastoreI interface {
	Load(args ...string) (users sql.Result ,err error)
	Save(name string, password string, email string) error
	Del(email string, password string ) error
}