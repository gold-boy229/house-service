package model

type AppUser struct {
	Id           int
	Role         string
	Email        string
	PasswordHash string
	UUID         string
}
