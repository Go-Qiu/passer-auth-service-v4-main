package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// pointer to a db struct (in go std package)
// that has all the data handling interfaces
// and database connections pool (managed by go via the database driver package).

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := `{
		"ok" : true,
		"msg" : "Reached GetAllUsers endpoint.",
		"data" : {}
	}`
	w.Write([]byte(body))
	//
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// setup the response header attributes
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// prepare the response body.
	body := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "User Id passed in : %s"
		"data" : {}
	}`, params["userid"])

	// write the content and send the response back to the requestor.
	w.Write([]byte(body))
	//
}

func AddUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

}
