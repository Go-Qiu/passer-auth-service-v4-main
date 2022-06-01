package routes

import (
	"log"
	"passer-auth-service-v4/pkg/controllers"

	"github.com/gorilla/mux"
)

func init() {
	log.Println("triggered init() at auth.routes.go")
}

// RegisterAuthRoutes sets all the route patterns and handlers maps
// for handling Authentication requests.
var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/auth", controllers.Auth).Methods("POST")
	router.HandleFunc("/verifytoken", controllers.VerifyToken).Methods("GET")
}
