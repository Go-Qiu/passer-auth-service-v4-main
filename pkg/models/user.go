package models

import (
	"database/sql"
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
}

func (u *User) CreateUser() *User {

	// create user code here.

	return u
}

func (u *User) GetAll(db *sql.DB) ([]User, error) {

	stmt := `SELECT 
	  id
	, email
	, nameFirst
	, nameLast
	, isActive
	, isAgent
	, dateCreated
	FROM PASSER_AUTH.Users
	`
	results, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	// ok.
	// slice to cache the results.
	var users []User

	for results.Next() {
		var user User
		err = results.Scan(&user.Id, &user.Email, &user.NameFirst, &user.NameLast, &user.IsActive, &user.IsAgent, &user.DateCreated)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
	//
}

func (u *User) GetUserById(db *sql.DB, id string) *User {

	// get user by id code here

	return u
}

func (u *User) DeleteUserById(db *sql.DB, id string) error {

	// delete user by id code here.

	return nil
}
