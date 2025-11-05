package handlers

import "github.com/labstack/echo/v4"

type userCreator interface {
	User_Create()
}

type registerHandler struct {
	repo *userCreator
}

func NewRegisterHandler(repo userCreator) *registerHandler {
	return &registerHandler{repo: &repo}
}

func (h *registerHandler) RegisterNewUser(c echo.Context) error {
	return nil
}
