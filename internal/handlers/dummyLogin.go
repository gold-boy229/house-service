package handlers

import "github.com/labstack/echo/v4"

type dummyLoginHandler struct {
}

func NewDummyLoginHandler() *dummyLoginHandler {
	return &dummyLoginHandler{}
}

func (h *dummyLoginHandler) DummyLogin(c echo.Context) error {
	return nil
}
