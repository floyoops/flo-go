package internal

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"html/template"
	"net/http"
	"os"
)

type App struct {
	echo *echo.Echo
}

func customHTTPErrorHandler(err error, c echo.Context) {
	log.Errorf(err.Error())
	errorHttp := echo.NewHTTPError(http.StatusInternalServerError)
	if errors.Is(err, repository.ErrOnSaveContact) {
		errorHttp.Message = "une erreur est survenue pendant la sauvegarde veuillez réessayer ultérieurement"
	} else if errors.Is(err, mailer.ErrOnSend) {
		errorHttp.Message = "une erreur est survenue pendant l envoie du mail veuillez réessayer ultérieurement"
	} else {
		errorHttp.Message = "une erreur est survenue veuillez réessayer ultérieurement"
	}
	err = c.JSON(http.StatusInternalServerError, errorHttp)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
}

func NewApp() *App {
	rootPath, _ := os.Getwd()
	renderer := &infra.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(rootPath + "/public/*.html")),
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.HTTPErrorHandler = customHTTPErrorHandler

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
