package viewmodel

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Router struct {
	View ViewI
}

func (r *Router) New() (router *httprouter.Router) {
	router = httprouter.New()
	router.POST("/auth/signin", r.View.SignIn)
	router.POST("/user/profile", r.View.Show)
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
