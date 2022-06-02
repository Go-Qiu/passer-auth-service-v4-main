package routes

import (
	"passer-auth-service-v4/pkg/controllers"

	"github.com/gorilla/mux"
)

// RegisterUsersRoutes sets all the route path patterns and handlers maps
// for handling Users CRUD API requests.
var RegisterUsersRoutes = func(router *mux.Router, ctl *controllers.CrudCtl) {

	router.HandleFunc("/", ctl.GetAll).Methods("GET")
	router.HandleFunc("/", ctl.Create).Methods("POST")
	router.HandleFunc("/email/{email}", ctl.GetByEmail).Methods("GET")
	router.HandleFunc("/{userid}", ctl.GetById).Methods("GET")
	router.HandleFunc("/{userid}", ctl.UpdateById).Methods("PUT")
	router.HandleFunc("/{userid}", ctl.DeleteById).Methods("DELETE")
}
