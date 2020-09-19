/*
Interface for Views.
This is important to ensure modularity.
do not edit or delete.
@author SashaCollins
*/
package viewmodel

import (
	"github.com/julienschmidt/httprouter"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"net/http"
)

type Response struct{
	Success 		bool                    	`json:"success"`
	Email 			string               		`json:"email"`
	Plugins 		[]data.Plugin           	`json:"plugins"`
	RepoData		map[string]interface{}		`json:"repos"`
	CourseData		map[string]interface{}		`json:"courses"`
}

type Request struct {
	Option 			string						`json:"option"`
	NewEmail 		string 						`json:"new_email"`
	Email 			string 						`json:"email"`
	Password 		string 						`json:"password"`
	Plugins 		[]data.Plugin 				`json:"plugins"`
	Repository 		string						`json:"repo"`
	Course 			string			 			`json:"course"`
}

type View interface {
	SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
	SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params)

	Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params)

	Repositories(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
	Courses(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
	Teams(w http.ResponseWriter, req *http.Request, ps httprouter.Params)
}