package app

import "github.com/labstack/echo/v4"

type houseHandler interface {
	GetHouseById(c echo.Context) error
	SubscribeForHouseUpdates(c echo.Context) error
	CreateNewHouse(c echo.Context) error
}

type flatHandler interface {
	Update(c echo.Context) error
	Create(c echo.Context) error
}

type dummyLoginHandler interface {
	DummyLogin(c echo.Context) error
}

type loginHandler interface {
	Login(c echo.Context) error
}

type registerHandler interface {
	RegisterNewUser(c echo.Context) error
}
