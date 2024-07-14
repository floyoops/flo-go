package http

import (
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
)

type EchoFactory struct {
	di                 *infra.Container
	allowOrigins       bool
	customErrorHandler bool
	templateRenderer   bool
	router             bool
}

func NewEchoFactory(container *infra.Container) *EchoFactory {
	return &EchoFactory{
		di:                 container,
		allowOrigins:       false,
		customErrorHandler: false,
		templateRenderer:   false,
		router:             false,
	}
}

func (f *EchoFactory) WithCors(value bool) *EchoFactory {
	f.allowOrigins = value
	return f
}

func (f *EchoFactory) WithCustomErrorHandler(value bool) *EchoFactory {
	f.customErrorHandler = value
	return f
}

func (f *EchoFactory) WithTemplateRenderer(value bool) *EchoFactory {
	f.templateRenderer = value
	return f
}

func (f *EchoFactory) WithRouter(value bool) *EchoFactory {
	f.router = value
	return f
}

func (f *EchoFactory) Build() *echo.Echo {
	e := echo.New()
	if f.allowOrigins {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: f.di.Config.HttpAllowOrigins,
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}
	if f.customErrorHandler {
		e.HTTPErrorHandler = customHTTPErrorHandler
	}
	if f.templateRenderer == true {
		renderer := &infra.TemplateRenderer{
			Templates: template.Must(template.ParseGlob(f.di.Config.RootPath + "/public/*.html")),
		}
		e.Renderer = renderer
	}
	if f.router == true {
		router(e, f.di)
	}

	return e
}
