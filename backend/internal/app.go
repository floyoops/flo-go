package internal

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/internal/infra/http"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

func NewApp(echoFactory *http.ServerFactory) (*App, error) {
	return &App{echo: echoFactory.Build()}, nil
}

func (a *App) Start(port int) error {
	if err := a.echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		return errors.New("Error on start server " + err.Error())
	}
	return nil
}
