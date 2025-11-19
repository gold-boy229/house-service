package dto

type RegisterUser_Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	UserType string `json:"user_type" validate:"required,oneof=client moderator"`
}

type RegisterUser_Response struct {
	UserUUID string `json:"user_id"`
}
