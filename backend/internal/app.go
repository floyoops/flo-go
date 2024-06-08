package internal

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"os"
)

type App struct {
	echo *echo.Echo
}

func NewApp() *App {
	rootPath, _ := os.Getwd()
	renderer := &infra.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(rootPath + "/public/*.html")),
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Renderer = renderer
	infra.NewRouter(e, infra.NewContainer(rootPath)).Build()
	return &App{echo: e}
}

func (a *App) Start(port int) error {
	if err := a.echo.Start(fmt.Sprintf(":%d", port)); err != nil {
		return errors.New("Error on start server " + err.Error())
	}
	return nil
}
