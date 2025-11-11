package app

import (
	"house-store/internal/handlers"
	"house-store/internal/repository"
	"house-store/internal/utilities/auth"

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

	err := auth.LoadJWTSecret()
	if err != nil {
		panic(err.Error())
	}

	app.echo.GET("/api/v1/house/{id}", houseHandler.GetHouseById)
	app.echo.POST("/api/v1/house/{id}/subscribe", houseHandler.SubscribeForHouseUpdates)
	app.echo.POST("/api/v1/house/create", houseHandler.CreateNewHouse)

	app.echo.POST("/api/v1/flat/create", flatHandler.Create)
	app.echo.POST("/api/v1/flat/update", flatHandler.Update)

	app.echo.GET("/api/v1/dummyLogin", dummyLoginHandler.DummyLogin)

	app.echo.POST("/api/v1/login", loginHandler.Login)

	app.echo.POST("/api/v1/register", registerHandler.RegisterNewUser)

	app.echo.Start(":8080")
}
