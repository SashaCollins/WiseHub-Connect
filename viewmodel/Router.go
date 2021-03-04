/*
Router delegates requests from frontend to backend and back
@author SashaCollins
@version 1.0
 */
package viewmodel

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/harlow/authtoken"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github/SashaCollins/Wisehub-Connect/model/data"
	"github/SashaCollins/Wisehub-Connect/model/plugins"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Router struct {
	Datastore data.DatastoreI
	View ViewI
}

var (
	PluginMap map[string]plugins.PluginI
	jwtKey = []byte("GyOqHEOUmbtYMADLxXG3rrinGbh535my")
	mail Mail
)

type Claims struct {
	Email string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Response struct{
	Success 		bool                    	`json:"success"`
	Email 			string               		`json:"email"`
	Plugins 		[]data.Plugin           	`json:"plugins"`
	Data			map[string]string			`json:"pluginData"`
	Token			string						`json:"token"`
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

/*
Loads email from authorized header token
 */
func (r *Router) loadEMailTokenHeader(w http.ResponseWriter, req *http.Request) (*Claims, error) {
	authToken, err := authtoken.FromRequest(req)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid token", 670)
		return nil, err
	}
	token, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid token", 670)
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok && !token.Valid {
		log.Println("token claims not valid")
		http.Error(w, "invalid token", 670)
		return nil, err
	}
	return claims, nil
}

//func (r *Router) loadEMailFromCookieToken(w http.ResponseWriter, req *http.Request) string {
//	c, err := req.Cookie("refresh")
//	log.Println(c)
//	log.Println(err)
//	if err != nil {
//		if err == http.ErrNoCookie {
//			w.WriteHeader(http.StatusUnauthorized)
//			return ""
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return ""
//	}
//	tknStr := c.Value
//	claims := &Claims{}
//	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
//		return jwtKey, nil
//	})
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			w.WriteHeader(http.StatusUnauthorized)
//			return ""
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return ""
//	}
//	if !tkn.Valid {
//		w.WriteHeader(http.StatusUnauthorized)
//		return ""
//	}
//	return claims.Email
//}

/*
Loads plugins via PluginLoader
 */
func (r *Router) LoadPlugins() map[string]plugins.PluginI {
	var loader PluginLoader
	pluginMap, err := loader.LoadAllPlugins()
	if err != nil {
		log.Println("Could not load plugins!")
	}
	return pluginMap
}

func (r *Router) GenerateEMailURL(request *http.Request, email, password, requestPath string) string {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	validateToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	validate, err := validateToken.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return ""
	}
	url := strings.ReplaceAll(request.URL.String(), requestPath, "/validate")
	return fmt.Sprintf("%v%v?token=%v", os.Getenv("TARGET_URL"), url, validate)
}

/*
Fetches credentials for the plugins from datastore
 */
func (r *Router) LoadPluginCredentials(userEmail string) (credentials map[string]plugins.Credentials) {
	dbUser, err := r.Datastore.Load(userEmail)
	if err != nil {
		log.Println("User not found!")
	}
	credentials = make(map[string]plugins.Credentials)
	for _, v := range dbUser[0].Plugins {
		if v.UsernameHost == "" || v.Token == "" {
			continue
		}
		credentials[v.PluginName] = plugins.Credentials{UserNameHost: v.UsernameHost, Token: v.Token}
	}
	return credentials
}

func (r *Router) SignUp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}

	var user Request
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}

	dbUser, _ := r.Datastore.Load(user.Email)
	if len(dbUser) == 0 {
		url := r.GenerateEMailURL(req, user.Email, user.Password, "/api/auth/signup")
		dataStruct := struct {
			Url string
		}{
			Url: url,
		}
		status, err := mail.SendEmailSMTP([]string{user.Email}, dataStruct, "validate.txt")
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid email", 668)
		}
		if status {
			log.Println("Email sent successfully using SMTP")
		}
		return
	}
	if dbUser[0].Email != "" || dbUser[0].Email == user.Email {
		log.Println(err)
		http.Error(w, "User already exists!", 666)
		return
	}
}

func (r *Router) Validate(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}

	claim, err := r.loadEMailTokenHeader(w, req)
	if err != nil {
		log.Println("invalid token")
		http.Error(w, "invalid token", 670)
		return
	}

	err = r.Datastore.Create(claim.Password, claim.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "User already exists!", 666)
		return
	}

	var response Response
	response.Success = true
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

func (r *Router) Refresh(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}

	claim, err := r.loadEMailTokenHeader(w, req)
	if err != nil {
		log.Println("invalid token")
		http.Error(w, "invalid token", 670)
		return
	}

	_, err = r.Datastore.Load(claim.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid token", 670)
		return
	}

	expirationTime := time.Now().Add(48 * time.Hour)
	claims := &Claims{
		Email: claim.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refresh, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh",
		Value:   refresh,
		Expires: expirationTime,
		SameSite: http.SameSiteLaxMode,
	})

	expirationTime = time.Now().Add(12 * time.Hour)
	claims = &Claims{
		Email: claim.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	access, err := accessToken.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response Response
	response.Success = true
	response.Token = access

	resp, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
	}
	_, _ = w.Write(resp)
}

func (r *Router) SignIn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	var user Request
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	dbUser, err := r.Datastore.Load(user.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid email or password", 667)
		return
	}

	var response Response
	if dbUser[0].Password == user.Password {
		response.Success = true

		expirationTime := time.Now().Add(48 * time.Hour)
		claims := &Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		refresh, err := refreshToken.SignedString(jwtKey)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "refresh",
			Value:   refresh,
			Expires: expirationTime,
			SameSite: http.SameSiteLaxMode,
		})

		expirationTime = time.Now().Add(12 * time.Hour)
		claims = &Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		access, err := accessToken.SignedString(jwtKey)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response.Token = access

		resp, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", 500)
		}
		_, _ = w.Write(resp)
	} else {
		log.Println(err)
		http.Error(w, "Invalid email or password", 667)
	}
	return
}

/*
Requests a user profile by email from the datastore
Returns the user profile or an error
 */
func (r *Router) Profile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}

	claim, err := r.loadEMailTokenHeader(w, req)
	if err != nil {
		log.Println("invalid token")
		http.Error(w, "invalid token", 670)
		return
	}

	dbUser, err := r.Datastore.Load(claim.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid email", 668)
		return
	}

	var response Response
	response.Success = true
	response.Plugins = dbUser[0].Plugins
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

/*
Requests new password
sending an email is not implemented in v1.0
is not used in v1.0
 */
//func (r *Router) Forgot(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	if (*req).Method == "OPTIONS" {
//	return
//	}
//	reqBody, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		log.Println(err)
//	}
//
//	var user Request
//	err = json.Unmarshal(reqBody, &user)
//	if err != nil {
//		log.Println(err)
//	}
//
//	var response Response
//	dbUser, err := r.Datastore.Load(user.Email)
//	if err != nil {
//		log.Println(err)
//	}
//
//	//TODO send email
//
//	response.Success = true
//	resp, err := json.Marshal(response)
//	if err != nil {
//		log.Println(err)
//		http.Error(w, "Internal server error", 500)
//		return
//	}
//	_, _ = w.Write(resp)
//	return
//}

/*
Requests an update for user profile from datastore for either email, password or credentials
Returns a success or error message
 */
func (r *Router) Update(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	var update Request
	err = json.Unmarshal(reqBody, &update)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	claim, err := r.loadEMailTokenHeader(w, req)
	if err != nil {
		log.Println("invalid token")
		http.Error(w, "invalid token", 670)
		return
	}

	var response Response
	switch update.Option {
	case "password":
		if _, err := r.Datastore.Load(claim.Email); err != nil {
			log.Println(err)
			http.Error(w, "Invalid email", 668)
			return
		}
		change := make(map[string]interface{})
		change["email"] = claim.Email
		change["password"] = update.Password
		if err = r.Datastore.Update(update.Option, change); err != nil {
			log.Println(err)
			http.Error(w, "Invalid email", 668)
			return
		}
	case "credentials":
		change := make(map[string]interface{})
		change["email"] = claim.Email
		change["updatedPlugins"] = update.Plugins
		if err := r.Datastore.Update(update.Option, change); err != nil {
			log.Println(err)
			http.Error(w, "Invalid email", 668)
			return
		}
	default:
		log.Println("Thanks for the fish!")
	}
	response.Success = true
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}
/*
Requests all data to show from a view
Returns fetched data
 */
func (r *Router) Show(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	var request Request
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	claim, err := r.loadEMailTokenHeader(w, req)
	if err != nil {
		log.Println("invalid token")
		http.Error(w, "invalid token", 670)
		return
	}

	_, err = r.Datastore.Load(claim.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid email", 668)
		return
	}

	credentials := r.LoadPluginCredentials(claim.Email)
	if len(credentials) == 0 {
		log.Println("invalid data")
		http.Error(w, "invalid data", 669)
		return
	}
	switch request.Option {
	case "general":
		r.View = &GeneralView{}
	case "template":
		//r.View = &TemplateView{}
		log.Println("Template")
	default:
		//r.View = &DefaultView{}
		log.Println("test")
	}
	r.View.SetPlugins(PluginMap)
	r.View.SetCredentials(credentials)
	pluginData, err := r.View.GetData()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 669)
		return
	}

	var response Response
	response.Success = true
	response.Data = pluginData
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

/*
Requests a deletion of a user profile from datastore
Returns a success or error message
not used in v1.0
 */
func (r *Router) Delete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if (*req).Method == "OPTIONS" {
		return
	}
	var response Response

	claim, err := r.loadEMailTokenHeader(w, req)
	if err != nil {
		log.Println("invalid token")
		http.Error(w, "invalid token", 670)
		return
	}

	if _, err := r.Datastore.Load(claim.Email); err != nil {
		log.Printf("DeleteProfile: %v\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}
	if err := r.Datastore.Delete(claim.Email); err != nil {
		log.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Invalid email", 668)
		return
	}
	response.Success = true
	resp, err := json.Marshal(response)
	fmt.Println(resp)
	if err != nil {
		log.Printf("DeleteProfile: %s\n", err)
		http.Error(w, "Internal server error", 500)
		return
	}
	_, _ = w.Write(resp)
	return
}

/*
Passes on the incoming http Requests
If functionality is extended add new routes here
 */
func (r *Router) New() (handler http.Handler) {
	router := httprouter.New()

	// Authentication
	router.POST("/api/auth/signin", r.SignIn)
	router.POST("/api/auth/signup", r.SignUp)
	router.POST("/api/auth/refresh", r.Refresh)
	router.GET("/api/auth/validate", r.Validate)
	//router.GET("/api/auth/forgot", r.Forgot)

	// Profile
	router.GET("/api/user/profile", r.Profile)
	router.POST("/api/user/update/password", r.Update)
	router.POST("/api/user/update/credentials", r.Update)
	//router.GET("/api/user/delete", r.View.Delete)

	// Fetch view data
	router.POST("/api/data/all", r.Show)

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Content-Type", "application/json")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	httpCors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPut},
		ExposedHeaders: []string{"Authorization"},
		AllowedHeaders: []string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		// Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler = httpCors.Handler(router)

	return
}
/*
Starts the router
router should be running in a go routine
 */
func (r *Router) Run(port int, finished chan bool) {
	router := r.New()
	PluginMap = r.LoadPlugins()
	log.Printf("Run: %s\n", http.ListenAndServe(fmt.Sprintf(":%d", port), router))
	finished <- true
}
