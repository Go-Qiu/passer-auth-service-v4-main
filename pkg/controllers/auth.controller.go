package controllers

import (
	"log"
	"net/http"
)

func init() {
	log.Println("triggered init() in auth.controller.go")
}

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body := `{
		"ok" : true,
		"msg" : "Reached Auth endpoint.",
		"data" : {}
	}`

	w.Write([]byte(body))
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body := `{
		"ok" : true,
		"msg" : "Reached VerifyToken endpoint.",
		"data" : {}
	}`

	w.Write([]byte(body))
}
