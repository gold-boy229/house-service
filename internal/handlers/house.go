package handlers

import "github.com/labstack/echo/v4"

type houseProvider interface {
	House_GetById()
	House_SubscribeForUpdates()
	House_Create()
}

type houseHandler struct {
	repo *houseProvider
}

func NewHouseHandler(repo houseProvider) *houseHandler {
	return &houseHandler{repo: &repo}
}

func (h *houseHandler) GetHouseById(c echo.Context) error {
	return nil
}

func (h *houseHandler) SubscribeForHouseUpdates(c echo.Context) error {
	return nil
}

func (h *houseHandler) CreateNewHouse(c echo.Context) error {
	return nil
}
