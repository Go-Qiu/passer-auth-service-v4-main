package routes

import (
	"passer-auth-service-v4/pkg/controllers"
	"passer-auth-service-v4/pkg/middlewares"

	"github.com/gorilla/mux"
)

// RegisterUsersRoutes sets all the route path patterns and handlers maps
// for handling Users CRUD API requests.
var RegisterUsersRoutes = func(router *mux.Router, ctl *controllers.CrudCtl) {

	router.HandleFunc("/", middlewares.ValidateToken(ctl.GetAll)).Methods("GET")
	router.HandleFunc("/", middlewares.ValidateToken(ctl.Create)).Methods("POST")
	router.HandleFunc("/email/{email}", middlewares.ValidateToken(ctl.GetByEmail)).Methods("GET")
	router.HandleFunc("/{id}/pw/", middlewares.ValidateToken(ctl.ResetPasswordById)).Methods("POST")
	router.HandleFunc("/{userid}", middlewares.ValidateToken(ctl.GetById)).Methods("GET")
	router.HandleFunc("/{id}", middlewares.ValidateToken(ctl.UpdateById)).Methods("PUT")
	router.HandleFunc("/{id}", middlewares.ValidateToken(ctl.DeleteById)).Methods("DELETE")
}
