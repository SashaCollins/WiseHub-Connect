package viewmodel

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Router struct {
	View View
}

func (r *Router) New() (router *httprouter.Router) {
	router = httprouter.New()
	router.POST("/auth/signup", r.View.SignUp)
	router.POST("/auth/signin", r.View.SignIn)
	router.POST("/user/profile", r.View.Show)
	router.POST("/user/delete", r.View.Delete)
	router.POST("/user/update/email", r.View.Update)
	router.POST("/user/update/password", r.View.Update)
	router.POST("/user/update/plugins", r.View.Update)
	router.POST("/user/repos", r.View.Repositories)
	router.POST("/user/all", r.View.Courses)
	router.POST("/user/teams", r.View.Teams)
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	return
}
