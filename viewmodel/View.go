/*
Interface for Views.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package viewmodel

import (
	"github/SashaCollins/Wisehub-Connect/model/data"
	_ "github/SashaCollins/Wisehub-Connect/model/listener"
)

type Response struct{
	Success bool `json:"success"`
	Email string `json:"email"`
	Plugins []data.Plugin `json:"plugins"`
}

type UpdateEmail struct {
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

type UpdatePlugins struct {
	Email string `json:"email"`
	Plugins []Plugin `json:"plugins"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Plugins  string `json:"plugins"`
}

type Plugin struct {
	Name string `json:"name"`
	Token string `json:"token"`
	Description string `json:"description"`
	Updated bool `json:"updated"`
}

type View interface {
	SignUp()
	SignIn()
	Show()
	Update()
}