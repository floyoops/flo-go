package internal

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/internal/infra/http"
	"github.com/floyoops/flo-go/backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo   *echo.Echo
	logger logger.Logger
}

func NewApp(echoFactory *http.ServerFactory, logger logger.Logger) (*App, error) {
	return &App{echo: echoFactory.Build(), logger: logger}, nil
}

func (a *App) Start(port int) error {
	a.logger.Info("Starting server")
	if err := a.echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		a.logger.Error("Error on start server " + err.Error())
		return errors.New("Error on start server " + err.Error())
	}
	a.logger.Info("Server started success")
	return nil
}
