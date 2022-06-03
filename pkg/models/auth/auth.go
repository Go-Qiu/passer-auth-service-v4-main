package modelsAuth

import (
	"database/sql"
	modelsUser "passer-auth-service-v4/pkg/models/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthModel struct {
	db *sql.DB
}

// AuthParams is a struct used for parsing the JSON body of a auth request.
type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(db *sql.DB) *AuthModel {
	return &AuthModel{db: db}
}

// Auth execute the check against the database, to retrieve the user data that matches the following conditions:
// - isActive is true;
// - email matches.
// The pwHas retrieved will be used to compare against the password (submitted) using the bcrypt CompareHashAndPassword function.
// Returns the following:
// - the user data (without the password hash). nil if error has occurred.
// - error (if any). nil if no error.
func (am *AuthModel) Auth(input AuthParams) (*modelsUser.User, error) {

	u := modelsUser.AuthUser{}
	user := modelsUser.User{}

	stmt := `SELECT 
		id
		, email
		, nameFirst
		, nameLast
		, isAgent
		, isActive
		, pwHash
	FROM Users 
	WHERE 
		isActive = ? AND
		email = ?
	;`

	row := am.db.QueryRow(stmt, true, input.Email)
	err := row.Scan(&u.Id, &u.Email, &u.NameFirst, &u.NameLast, &u.IsAgent, &u.IsActive, &u.PwHash)
	if err != nil {
		return nil, err
	}

	// bcrypt compare
	err = bcrypt.CompareHashAndPassword([]byte(u.PwHash), []byte(input.Password))
	if err != nil {
		// pwhash does not match.
		return nil, err
	}

	// ok. authenticated.
	// stripe out the pwhash
	user.Id = u.Id
	user.Email = u.Email
	user.NameFirst = u.NameFirst
	user.NameLast = u.NameLast
	user.IsAgent = u.IsAgent
	user.IsActive = u.IsActive

	return &user, nil
}
