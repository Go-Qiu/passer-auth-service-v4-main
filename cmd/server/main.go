package main

import (
	"log"
	"net/http"
	"os"

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

	// declare and instantiate a web application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	r := app.routesV2()

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
