package models

import (
	"database/sql"
	"log"
	"passer-auth-service-v4/pkg/config"
)

var db *sql.DB

type User struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	NameFirst   string `json:"nameFirst"`
	NameLast    string `json:"nameLast"`
	IsActive    bool   `json:"isActive"`
	IsAgent     bool   `json:"isAgent"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func init() {
	log.Println("triggered init() at user.go")
	config.Connect()
	db = config.GetDB()
}

func (u *User) CreateUser() *User {

	// create user code here.

	return u
}

func (u *User) GetAllUsers() []User {
	var Users []User

	// get all users code here

	return Users
}

func (u *User) GetUserById(id string) *User {

	// get user by id code here

	return u
}

func (u *User) DeleteUserById(id string) error {

	// delete user by id code here.

	return nil
}
