package router

import (
	"github.com/gorilla/mux"
	controller "github.com/sachinchaudhary003/golangAuth/Controller"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/signup", controller.UserSignup).Methods("POST")
	r.HandleFunc("/login", controller.Userlogin).Methods("GET")

	return r

}
