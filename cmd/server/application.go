package main

import (
	"log"
	"net/http"
	"passer-auth-service-v4/pkg/controllers"
	"passer-auth-service-v4/pkg/routes"

	"github.com/gorilla/mux"
)

// application struct is for facilitating the implementation of the dependencies injection model.
type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	crudCtlUsers *controllers.CrudCtl
	authCtl      *controllers.AuthCtl
}

// Users method to direct request to operate on users related data to
// the appropriate user data operations handler.
func (a *application) Users(w http.ResponseWriter, r *http.Request) {

	// users.Handler(w, r, a.dataStore)

}

// routes returns a gorilla/mux router that contains all the route-handler mappings.
func (a *application) routes() *mux.Router {

	// instantiate a gorilla mux router.
	r := mux.NewRouter()

	// register all the defined routes (for users requests) with
	// the new mux router instance.
	g := r.PathPrefix("/api/v4").Subrouter()

	// authentication routes
	routes.RegisterAuthRoutes(g, a.authCtl)

	// users routes
	userRoutes := g.PathPrefix("/users/").Subrouter()
	routes.RegisterUsersRoutes(userRoutes, a.crudCtlUsers)

	return r
}
