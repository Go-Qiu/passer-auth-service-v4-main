package controllers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	modelsUser "passer-auth-service-v4/pkg/models/user"
	"passer-auth-service-v4/pkg/utils"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// CrudCtl is a struct that represents a CRUD controller, in a MVC pattern.
type CrudCtl struct {
	db   *sql.DB
	name string
}

// paramsAdd is the struct for storing the data
// passed in via the add user request body.
type paramsAdd struct {
	Email     string `json:"email"`
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Password  string `json:"password"`
	IsActive  bool   `json:"isActive"`
	IsAgent   bool   `json:"isAgent"`
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

	um := modelsUser.New(ctl.db)
	u, err := um.GetAll()
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
	rows := len(*u)
	var msg string

	if rows > 1 {
		// pural form
		msg = fmt.Sprintf(`%d rows of user data found.`, rows)
	} else {
		// singular form
		msg = fmt.Sprintf(`%d row of user data found.`, rows)
	}

	utils.SendDataToClient(&w, list, msg)
	//
}

// GetById return the specific user record
func (ctl *CrudCtl) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// setup the response header attributes
	w.Header().Set("Content-Type", "application/json")

	um := modelsUser.New(ctl.db)
	id := params["userid"]
	if id == "" {
		// empty id string
		customErr := errors.New(`[USER-CTL] userid is a require input and cannot be empty`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	// sanitize the id here.

	// ok.
	u, err := um.GetUserById(id)
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
	utils.SendDataToClient(&w, data, "user data found")
	//
}

// GetByEmail return the specific user record
func (ctl *CrudCtl) GetByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// setup the response header attributes
	w.Header().Set("Content-Type", "application/json")

	um := modelsUser.New(ctl.db)
	emailB64 := params["email"]
	if emailB64 == "" {
		// empty id string
		customErr := errors.New(`[USER-CTL] email is a require input and cannot be empty`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// decode the base64 encoded email parameter.
	emailByte, err := base64.StdEncoding.DecodeString(emailB64)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to read email attribute`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}
	email := string(emailByte)

	// sanitize the id here.

	// ok.
	u, err := um.GetUserByEmail(email)
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
	utils.SendDataToClient(&w, data, "user data found")
	//
}

func (ctl *CrudCtl) Create(w http.ResponseWriter, r *http.Request) {

	// read the setting from .env
	err := godotenv.Load("../../.env")
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to read .env file`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}
	dateFormatString := os.Getenv("DATE_CREATED_FORMAT")

	w.Header().Set("Content-Type", "application/json")

	newUserIn := paramsAdd{}
	newUser := modelsUser.NewUserRegister{}

	// instantiate a new UserModel struct.
	um := modelsUser.New(ctl.db)

	err = utils.ParseBody(r, &newUserIn)
	if err != nil {
		// fail to read json body.
		customErr := errors.New(`[USER-CTL] fail to read JSON body`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// checks.  minimize sql injection risk.
	ok, err := sanitizeCreateInput(newUserIn)
	if !ok || err != nil {
		// fail to sanitize
		errmsg := fmt.Sprintf(`[USER-CTL] %s`, err)
		customErr := errors.New(errmsg)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// bcrypt the password.
	pwhash, err := bcrypt.GenerateFromPassword([]byte(newUserIn.Password), bcrypt.MinCost)
	if err != nil {
		// fail to hash password.
		customErr := errors.New(`[USER-CTL] fail to hash password`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}
	newUser.PwHash = string(pwhash)

	// generate user id.
	newUser.Id = utils.GenerateID()
	newUser.Email = newUserIn.Email
	newUser.NameFirst = newUserIn.NameFirst
	newUser.NameLast = newUserIn.NameLast
	newUser.IsActive = newUserIn.IsActive
	newUser.IsAgent = newUserIn.IsAgent

	// created date.
	newUser.DateCreated = time.Now().Local().Format(dateFormatString)

	// add the new user into database
	err = um.AddUser(&newUser)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to add user`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// get the record added from the database
	u, err := um.GetUserById(newUser.Id)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to find the newly added data`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// marshal the record into JSON.
	data, err := json.Marshal(u)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to parse the added data into JSON`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// ok.
	utils.SendDataToClient(&w, data, "user data added")
	//
}

func (ctl *CrudCtl) UpdateById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	if strings.TrimSpace(id) == "" {
		// empty id string
		customErr := errors.New(`[USER-CTL] id is a require input and cannot be empty`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// get the specific user

	// get the JSON body
	update := modelsUser.UpdateUser{}
	err := utils.ParseBody(r, &update)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to parse result into JSON`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	// checks. minimize sql injection risk.
	ok, err := sanitizeUpdateInput(update)
	if !ok || err != nil {
		// fail to sanitize
		errmsg := fmt.Sprintf(`[USER-CTL] %s`, err)
		customErr := errors.New(errmsg)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// instantiate a user model struct.
	um := modelsUser.New(ctl.db)

	// check if the user data exist (before updating).
	_, err = um.GetUserById(id)
	if err != nil {
		// not found.
		customErr := errors.New(`[USER-CTL] user data not found`)
		utils.SendNotFoundMsgToClient(&w, customErr)
		return
	}

	err = um.UpdateUserById(id, update)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to update user data`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// get the user data (after update).
	up, err := um.GetUserById(id)
	if err != nil {
		// not found.
		customErr := errors.New(`[USER-CTL] user data not found`)
		utils.SendNotFoundMsgToClient(&w, customErr)
		return
	}

	data, err := json.Marshal(up)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to parse JSON data`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// prepare the response.
	utils.SendDataToClient(&w, data, `[USER-CTL] user data updated`)

}

func (ctl *CrudCtl) DeleteById(w http.ResponseWriter, r *http.Request) {

	// setup the response header.
	w.Header().Set("Content-Type", "application/json")

	// get the id from the url
	params := mux.Vars(r)
	id := params["id"]

	// instantiate a user model.
	um := modelsUser.New(ctl.db)

	// check for empty id
	if strings.TrimSpace(id) == "" {
		// empty id
		customErr := errors.New(`[USER-CTL] required id is not passed in or empty`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	// check if the id is of the specific format

	// check if the user data (to be deleted) exist.
	u, err := um.GetUserById(id)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to find the user data`)
		utils.SendNotFoundMsgToClient(&w, customErr)
		return
	}

	// marshal the record (to be deleted) into JSON.
	data, err := json.Marshal(u)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to parse the user data into JSON`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	err = um.DeleteUserById(id)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to delete user`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	utils.SendDataToClient(&w, data, "user data deleted")

}

// ResetPasswordById reset the password of the specific user, identified by id.
func (ctl *CrudCtl) ResetPasswordById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	if strings.TrimSpace(id) == "" {
		// empty id string
		customErr := errors.New(`[USER-CTL] id is a require input and cannot be empty`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// get the JSON body.
	reset := modelsUser.ResetPasswordParams{}
	err := utils.ParseBody(r, &reset)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to parse result into JSON`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	// checks. minimize sql injection risk.
	ok, err := sanitizePasswordResetInput(reset)
	if !ok || err != nil {
		// fail to sanitize
		errmsg := fmt.Sprintf(`[USER-CTL] %s`, err)
		customErr := errors.New(errmsg)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// instantiate a user model struct.
	um := modelsUser.New(ctl.db)

	// check if the user data exist (before updating).
	_, err = um.GetUserById(id)
	if err != nil {
		// not found.
		customErr := errors.New(`[USER-CTL] user data not found`)
		utils.SendNotFoundMsgToClient(&w, customErr)
		return
	}

	// bcrypt the password.
	pwhash, err := bcrypt.GenerateFromPassword([]byte(reset.Password), bcrypt.MinCost)
	if err != nil {
		// fail to hash password.
		customErr := errors.New(`[USER-CTL] fail to process the password`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// prepare the parameter for password reset
	// using the user model
	resetHash := modelsUser.ResetPasswordParams{}
	resetHash.Password = string(pwhash)

	err = um.ResetPassword(id, resetHash)
	if err != nil {
		customErr := errors.New(`[USER-CTL] fail to reset password`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// prepare the response.
	utils.SendDataToClient(&w, nil, `[USER-CTL] password reset done`)

}

// sanitizeCreateInput executes checks all the input values passed in for creating user data.
func sanitizeCreateInput(input paramsAdd) (bool, error) {

	// check for empty string.
	if strings.TrimSpace(input.Password) == "" {
		// password is empty.
		return false, errors.New(`required attribute, "password" cannot be an empty value`)
	}

	if strings.TrimSpace(input.NameFirst) == "" {
		// first name is empty.
		return false, errors.New(`required attribute, "nameFirst" cannot be an empty value`)
	}

	if strings.TrimSpace(input.NameLast) == "" {
		// last name is empty.
		return false, errors.New(`required attribute, "nameLast" cannot be an empty value`)
	}

	if strings.TrimSpace(input.Email) == "" {
		// last name is empty.
		return false, errors.New(`required attribute, "email" cannot be an empty value`)
	}

	// email formate check
	if !utils.IsValidEmailFormat(input.Email) {
		return false, errors.New(`required attribute, "email" must be in a valid format`)
	}

	// alpha and space only check
	if !utils.IsAlphaAndSpaceOnly(input.NameFirst) {
		// fist name contains not allowed characters
		return false, errors.New(`required attribute, "nameFirst" should contain only alpahbets or space only`)
	}

	if !utils.IsAlphaAndSpaceOnly(input.NameLast) {
		// fist name contains not allowed characters
		return false, errors.New(`required attribute, "nameLast" should contain only alpahbets or space only`)
	}

	// allowed characters only check
	if !utils.AreAllowedCharacters(input.Password) {
		// password contains not allowed characters
		return false, errors.New(`required attribute, "password" should contain characters that are not allowed`)
	}

	return true, nil
}

// sanitizeUpdateInput execute checks on the input parameters passed in for user data update.
func sanitizeUpdateInput(input modelsUser.UpdateUser) (bool, error) {

	// check for empty string.

	if strings.TrimSpace(input.NameFirst) == "" {
		// first name is empty.
		return false, errors.New(`required attribute, "nameFirst" cannot be an empty value`)
	}

	if strings.TrimSpace(input.NameLast) == "" {
		// last name is empty.
		return false, errors.New(`required attribute, "nameLast" cannot be an empty value`)
	}

	// alpha and space only check
	if !utils.IsAlphaAndSpaceOnly(input.NameFirst) {
		// fist name contains not allowed characters
		return false, errors.New(`required attribute, "nameFirst" should contain only alpahbets or space only`)
	}

	if !utils.IsAlphaAndSpaceOnly(input.NameLast) {
		// fist name contains not allowed characters
		return false, errors.New(`required attribute, "nameLast" should contain only alpahbets or space only`)
	}

	return true, nil
}

// sanitizePasswordResetInput execute checks on the input parameter passed in for password reset.
func sanitizePasswordResetInput(input modelsUser.ResetPasswordParams) (bool, error) {

	// check for empty string.
	if strings.TrimSpace(input.Password) == "" {
		// password is empty.
		return false, errors.New(`required attribute, "password" cannot be an empty value`)
	}

	// allowed characters only check
	if !utils.AreAllowedCharacters(input.Password) {
		// password contains not allowed characters
		return false, errors.New(`required attribute, "password" should contain characters that are not allowed`)
	}

	return true, nil
}
