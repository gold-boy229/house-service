package dto

type Login_Request struct {
	UserUUID string `json:"id" validate:"required,uuid"`
	Password string `json:"password" validate:"required"`
}

type Login_Response struct {
	Token string `json:"token"`
}
