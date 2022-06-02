package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"passer-auth-service-v4/pkg/models"
	"passer-auth-service-v4/pkg/utils"

	"github.com/gorilla/mux"
)

// CrudCtl is a struct that represents a CRUD controller, in a MVC pattern.
type CrudCtl struct {
	db   *sql.DB
	name string
}

// New returns an instance of the CRUD controller.
func New(db *sql.DB, name string) *CrudCtl {
	ctl := &CrudCtl{db: db, name: name}
	return ctl
}

// GetAll returns all the User records
func (ctl *CrudCtl) GetAll(w http.ResponseWriter, r *http.Request) {

	// prepare the response header.
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	u, err := user.GetAll(ctl.db)
	if err != nil {
		// exception handling
		utils.SendErrorMsgToClient(&w, err)
		return
	}

	// marshal the records into JSON.
	list, err := json.Marshal(u)
	if err != nil {
		customErr := errors.New(`[USERS-CTL] fail to parse results into JSON`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// ok.
	utils.SendDataToClient(&w, list)
	//
}

// GetById return the specific user record
func (ctl *CrudCtl) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// setup the response header attributes
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	id := params["userid"]
	if id == "" {
		// empty id string
		customErr := errors.New(`[USER-CTL] userid is a require input and cannot be empty`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// sanitize the id here.

	// ok.
	u, err := user.GetUserById(ctl.db, id)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to execute the record search`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// marshal the record into JSON.
	data, err := json.Marshal(u)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to parse result into JSON`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// ok.
	utils.SendDataToClient(&w, data)
	//
}

func (ctl *CrudCtl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := `{
		"ok" : true,
		"msg" : "Reached Create endpoint.",
		"data" : {}
	}`
	w.Write([]byte(body))
}

func (ctl *CrudCtl) UpdateById(w http.ResponseWriter, r *http.Request) {

	// get the specific user

	// delete the specific user

	// add the specific user

	// prepare the response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := `{
		"ok" : true,
		"msg" : "Reached UpdateById endpoint.",
		"data" : {}
	}`

	// send out the response.
	w.Write([]byte(body))
}

func (ctl *CrudCtl) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := `{
		"ok" : true,
		"msg" : "Reached DeleteById endpoint.",
		"data" : {}
	}`
	w.Write([]byte(body))
}
