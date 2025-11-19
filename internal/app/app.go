package app

import (
	"fmt"
	"house-store/internal/config"
	"house-store/internal/handlers"
	mw "house-store/internal/middleware"
	"house-store/internal/repository"
	"house-store/internal/utilities/auth"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Return the error so Echo's error handler can process it
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewApp() *App {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
	return &App{echo: e}
}

func (app *App) Run() {
	configDB, err := config.ReadConfigDB()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Config loaded successfully\n ConfigDB = %+v\n", configDB)

	repo, err := repository.New(configDB)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Repository created successfully")

	err = auth.LoadJWTSecret()
	if err != nil {
		panic(err.Error())
	}

	var (
		houseHandler      houseHandler      = handlers.NewHouseHandler(repo)
		flatHandler       flatHandler       = handlers.NewFlatHandler(repo)
		dummyLoginHandler dummyLoginHandler = handlers.NewDummyLoginHandler()
		loginHandler      loginHandler      = handlers.NewLoginHandler(repo)
		registerHandler   registerHandler   = handlers.NewRegisterHandler(repo)
	)

	app.echo.GET("/api/v1/house/:id", houseHandler.GetHouseById, mw.AuthOnly)
	app.echo.POST("/api/v1/house/:id/subscribe", houseHandler.SubscribeForHouseUpdates, mw.AuthOnly)
	app.echo.POST("/api/v1/house/create", houseHandler.CreateNewHouse, mw.ModeratorsOnly)

	app.echo.POST("/api/v1/flat/create", flatHandler.Create, mw.AuthOnly)
	app.echo.PATCH("/api/v1/flat/update", flatHandler.Update, mw.ModeratorsOnly)

	app.echo.GET("/api/v1/dummyLogin", dummyLoginHandler.DummyLogin)

	app.echo.POST("/api/v1/login", loginHandler.Login)

	app.echo.POST("/api/v1/register", registerHandler.RegisterNewUser)

	app.echo.Start(":8080")
}
