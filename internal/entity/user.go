package entity

type User struct {
	UUID         string
	Role         string
	Email        string
	PasswordHash string
}

type UserLoginData struct {
	UUID     string
	Password string
}
type UserExistence struct {
	Exists            bool
	IsPasswordCorrect bool
	Role              string
}
