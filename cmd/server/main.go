package main

import (
	"log"
	"net/http"
	"os"
	"passer-auth-service-v4/pkg/controllers"

	"github.com/joho/godotenv"
)

func main() {

	// declare custom loggers
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// get .env values
	err := godotenv.Load("../../.env")
	if err != nil {
		errString := "[AUTH-SVC]: fail to load .env"
		errorLog.Fatal(errString)
	}
	ADDR := os.Getenv("SERVER_ADDR")
	DSN_AUTH := os.Getenv("DSN_AUTH")

	// instantiate a mysql connections pool struct;
	// and checked that the connection to the database is working.
	db, err := OpenDB(DSN_AUTH)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// instantiate a Users CRUD controller.
	crudUsers := controllers.New(db, "Users")

	// declare and instantiate a web application
	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		crudCtlUsers: crudUsers,
	}

	r := app.routes()

	// // associate the mux router with the root route of the http server
	// http.Handle("/", r)

	// declare and instantiate a custom http server.  This is a minor pre-cautionary step to minimize the cyber-risk exposure surface of this service.
	srv := http.Server{
		Addr:     ADDR,
		ErrorLog: app.errorLog,
		Handler:  r,
	}

	app.infoLog.Printf("HTTPS Server started and listening on https://%s ...", ADDR)
	err = srv.ListenAndServeTLS("../../ssl/cert03.pem", "../../ssl/key03.pem")
	if err != nil {
		app.errorLog.Fatal(err)
	}
	// log.Fatal(http.ListenAndServe(":9090", r))
}
