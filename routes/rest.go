package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/miia_api/controllers"
	"github.com/joaopandolfi/miia_api/mhandlers"
)

func rest(r *mux.Router) {
	//Common
	r.HandleFunc("/rest/login", controllers.AuthController{}.Login).Methods("POST")
	r.HandleFunc("/rest/logout", mhandlers.AuthProtection(controllers.AuthController{}.Logout)).Methods("POST", "GET")
	r.HandleFunc("/rest/check/auth", controllers.AuthController{}.CheckAuth).Methods("GET")

}
