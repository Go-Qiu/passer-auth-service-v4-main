package main

import (
	"fmt"
	"log"
	"net/http"
	"passer-auth-service-v4/pkg/controllers"
	"passer-auth-service-v4/pkg/routes"
	"runtime/debug"

	"github.com/gorilla/mux"
)

// application struct is for facilitating the implementation of the dependencies injection model.
type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	crudCtlUsers *controllers.CrudCtl
}

// Users method to direct request to operate on users related data to
// the appropriate user data operations handler.
func (a *application) Users(w http.ResponseWriter, r *http.Request) {

	// users.Handler(w, r, a.dataStore)

}

// serverError will log server side errors and send a HTTP Internal Server Error to the requestor.
func (a *application) serverError(w http.ResponseWriter, err error) {
	// log the error on the server side
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.errorLog.Println(trace)

	// send an http error response to the requestor.
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific http error, with status code and text, to the requestor.
func (a *application) clientError(w http.ResponseWriter, status int, msg ...string) {
	http.Error(w, http.StatusText(status), status)
}

// routes returns a server mux, containing all the path patterns to handlers mapping.
// func (a *application) routes() *http.ServeMux {

// 	// it is recommended not to  use the default server mux implementation in the http package, in production.
// 	// recommended to declare a custom server mux, for use in instantiating a http server, in production.
// 	mux := http.NewServeMux()

// 	// fixed path patterns
// 	mux.HandleFunc("/auth", a.Auth)
// 	mux.Handle("/users", middlewares.ValidateJWT(http.HandlerFunc(a.Users)))
// 	mux.Handle("/verify", middlewares.ValidateJWT(http.HandlerFunc(a.Verify)))
// 	return mux
// }

// routesv2 returns a gorilla/mux router that contains all the route-handler mappings.
func (a *application) routes() *mux.Router {

	// instantiate a gorilla mux router.
	r := mux.NewRouter()

	// register all the defined routes (for users requests) with
	// the new mux router instance.
	g := r.PathPrefix("/api/v4").Subrouter()

	// authentication routes
	routes.RegisterAuthRoutes(g)

	// users routes
	userRoutes := g.PathPrefix("/users/").Subrouter()
	routes.RegisterUsersRoutes(userRoutes, a.crudCtlUsers)

	return r
}
