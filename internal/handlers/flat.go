package handlers

import "github.com/labstack/echo/v4"

type flatProvider interface {
	Flat_Create()
	Flat_Update()
}

type flatHandler struct {
	repo *flatProvider
}

func NewFlatHandler(repo flatProvider) *flatHandler {
	return &flatHandler{repo: &repo}
}

func (h *flatHandler) Update(c echo.Context) error {
	return nil
}

func (h *flatHandler) Create(c echo.Context) error {
	return nil
}
