package modelsUser

import (
	"database/sql"
)

// var db *sql.DB

type User struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	NameFirst   string `json:"nameFirst"`
	NameLast    string `json:"nameLast"`
	IsActive    bool   `json:"isActive"`
	IsAgent     bool   `json:"isAgent"`
	DateCreated string `json:"dateCreated"`
}

type NewUserRegister struct {
	User
	PwHash string `json:"pwHash"`
}

type UserModel struct {
	db *sql.DB
}

func New(db *sql.DB) *UserModel {
	model := UserModel{db: db}
	return &model
}

func (um *UserModel) CreateUser() *User {

	// create user code here.
	u := &User{}
	return u
}

func (um *UserModel) GetAll() (*[]User, error) {

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
	results, err := um.db.Query(stmt)
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
	return &users, nil
	//
}

func (um *UserModel) GetUserById(id string) (*User, error) {

	// instantiate a User struct
	u := User{}

	stmt := `SELECT 
	id
	, email
	, nameFirst
	, nameLast
	, isActive
	, isAgent
	, dateCreated
	FROM PASSER_AUTH.Users 
	WHERE id = ?
	`
	row := um.db.QueryRow(stmt, id)
	err := row.Scan(&u.Id, &u.Email, &u.NameFirst, &u.NameLast, &u.IsActive, &u.IsAgent, &u.DateCreated)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (um *UserModel) GetUserByEmail(email string) (*User, error) {

	// instantiate a User struct.
	u := User{}

	stmt := `SELECT 
	id
	, email
	, nameFirst
	, nameLast
	, isActive
	, isAgent
	, dateCreated
	FROM PASSER_AUTH.Users 
	WHERE email = ?
	`
	row := um.db.QueryRow(stmt, email)
	err := row.Scan(&u.Id, &u.Email, &u.NameFirst, &u.NameLast, &u.IsActive, &u.IsAgent, &u.DateCreated)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Add will add a user record into the database.
func (um *UserModel) AddUser(new *NewUserRegister) error {

	stmt := `INSERT INTO Users (
		id, email
		, nameFirst, nameLast
		, isActive, isAgent
		, dateCreated, pwHash
	) VALUES (
		?, ?
		,?, ?
		,?, ?
		,?, ?
	);`

	// execute the prepared sql script.
	_, err := um.db.Exec(stmt, new.Id, new.Email, new.NameFirst, new.NameLast, new.IsActive, new.IsAgent, new.DateCreated, new.PwHash)
	if err != nil {
		return err
	}

	// ok. added.
	return nil
}

func (um *UserModel) DeleteUserById(id string) error {

	stmt := `DELETE FROM Users WHERE id = ?;`

	// execute the prepared sql script.
	_, err := um.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	// ok. deleted.
	return nil
}
