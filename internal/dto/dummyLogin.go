package dto

type DummyLoginRequest struct {
	UserType string `query:"user_type"`
}

type DummyLoginResponse struct {
	Token string `json:"token"`
}
