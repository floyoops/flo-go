package internal

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/floyoops/flo-go/backend/internal/infra/http"
	"github.com/labstack/echo/v4"
	"os"
)

type App struct {
	echo *echo.Echo
}

func NewApp() (*App, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	container := infra.NewContainer(rootPath)

	echoFactory := http.NewEchoFactory(container)
	echoFactory.WithCors([]string{"http://localhost:5173"})
	echoFactory.WithTemplateRenderer(true)
	echoFactory.WithCustomErrorHandler(true)
	echoFactory.WithRouter(true)

	return &App{echo: echoFactory.Build()}, nil
}

func (a *App) Start(port int) error {
	if err := a.echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		return errors.New("Error on start server " + err.Error())
	}
	return nil
}
