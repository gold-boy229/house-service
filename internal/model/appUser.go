package model

type AppUser struct {
	UUID         string
	Role         string
	Email        string
	PasswordHash string
}
