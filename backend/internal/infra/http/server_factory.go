package http

import (
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
)

type ServerFactory struct {
	rootPath         string
	httpAllowOrigins []string
	routes           []Route
}

func NewServerFactory(
	rootPath string,
	httpAllowOrigins []string,
	routes []Route,
) *ServerFactory {
	return &ServerFactory{
		rootPath:         rootPath,
		httpAllowOrigins: httpAllowOrigins,
		routes:           routes,
	}
}

func (f *ServerFactory) Build() *echo.Echo {
	e := echo.New()
	if len(f.httpAllowOrigins) > 0 {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: f.httpAllowOrigins,
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}
	e.HTTPErrorHandler = customHTTPErrorHandler

	renderer := &infra.TemplateRenderer{
		Templates: template.Must(template.ParseGlob(f.rootPath + "/public/*.html")),
	}
	e.Renderer = renderer

	for _, r := range f.routes {
		e.Add(r.method, r.path, r.handler)
	}

	return e
}
