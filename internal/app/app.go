package app

import (
	"house-store/internal/handlers"
	"house-store/internal/repository"

	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

func NewApp() *App {
	return &App{echo: echo.New()}
}

func (app *App) Run() {
	repo := repository.NewRepository()
	var (
		houseHandler      houseHandler      = handlers.NewHouseHandler(repo)
		flatHandler       flatHandler       = handlers.NewFlatHandler(repo)
		dummyLoginHandler dummyLoginHandler = handlers.NewDummyLoginHandler()
		loginHandler      loginHandler      = handlers.NewLoginHandler(repo)
		registerHandler   registerHandler   = handlers.NewRegisterHandler(repo)
	)

	app.echo.GET("/house/{id}", houseHandler.GetHouseById)
	app.echo.POST("/house/{id}/subscribe", houseHandler.SubscribeForHouseUpdates)
	app.echo.POST("/house/create", houseHandler.CreateNewHouse)

	app.echo.POST("/flat/create", flatHandler.Create)
	app.echo.POST("/flat/update", flatHandler.Update)

	app.echo.POST("/dummyLogin", dummyLoginHandler.DummyLogin)

	app.echo.POST("/login", loginHandler.Login)

	app.echo.POST("/register", registerHandler.RegisterNewUser)
}
