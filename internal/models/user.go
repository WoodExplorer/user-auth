package models

type User struct {
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

type UserIdentity struct {
	Name string `json:"name"`
}
