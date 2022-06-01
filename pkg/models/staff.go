package models

type Staff struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	NameFirst   string `json:"nameFirst"`
	NameLast    string `json:"nameLast"`
	IsActive    bool   `json:"isActive"`
	IsAdmin     bool   `json:"isAdmin"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}
