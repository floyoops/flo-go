package app

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/pkg/app/di"
	"github.com/floyoops/flo-go/pkg/app/render"
	"github.com/floyoops/flo-go/pkg/app/router"
	"github.com/labstack/echo/v4"
	"html/template"
	"os"
)

type App struct {
	echo *echo.Echo
}

func NewApp() *App {
	rootPath, _ := os.Getwd()
	renderer := &render.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(rootPath + "/public/*.html")),
	}
	e := echo.New()
	e.Renderer = renderer
	router.NewRouter(e, di.NewContainer()).Build()
	return &App{echo: e}
}

func (a *App) Start(port int) error {
	if err := a.echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		return errors.New("Error on start server " + err.Error())
	}
	return nil
}
