package routes

import (
	"passer-auth-service-v4/pkg/controllers"

	"github.com/gorilla/mux"
)

// RegisterUsersRoutes sets all the route path patterns and handlers maps
// for handling Users CRUD API requests.
var RegisterUsersRoutes = func(router *mux.Router) {

	router.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/", controllers.AddUser).Methods("POST")
	router.HandleFunc("/{userid}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/{userid}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/{userid}", controllers.DeleteUserById).Methods("DELETE")
}
