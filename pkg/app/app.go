package app

import (
	"errors"
	"fmt"
	appInfra "github.com/floyoops/flo-go/pkg/app/internal/infra"
	"github.com/labstack/echo/v4"
	"html/template"
	"os"
)

type App struct {
	echo *echo.Echo
}

func NewApp() *App {
	rootPath, _ := os.Getwd()
	renderer := &appInfra.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(rootPath + "/public/*.html")),
	}
	e := echo.New()
	e.Renderer = renderer
	appInfra.NewRouter(e, appInfra.NewContainer(rootPath)).Build()
	return &App{echo: e}
}

func (a *App) Start(port int) error {
	if err := a.echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		return errors.New("Error on start server " + err.Error())
	}
	return nil
}
