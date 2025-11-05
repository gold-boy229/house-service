package handlers

import "github.com/labstack/echo/v4"

type userProvider interface {
	User_Get()
}

type loginHandler struct {
	repo *userProvider
}

func NewLoginHandler(repo userProvider) *loginHandler {
	return &loginHandler{repo: &repo}
}

func (h *loginHandler) Login(c echo.Context) error {
	return nil
}
