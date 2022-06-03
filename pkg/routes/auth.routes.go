package routes

import (
	"log"
	"passer-auth-service-v4/pkg/controllers"
	"passer-auth-service-v4/pkg/middlewares"

	"github.com/gorilla/mux"
)

func init() {
	log.Println("triggered init() at auth.routes.go")
}

// RegisterAuthRoutes sets all the route patterns and handlers maps
// for handling Authentication requests.
var RegisterAuthRoutes = func(router *mux.Router, ctl *controllers.AuthCtl) {
	router.HandleFunc("/auth", ctl.Auth).Methods("POST")
	router.HandleFunc("/verifytoken", middlewares.ValidateToken(ctl.VerifyToken)).Methods("GET")
}
