package http

import (
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact"
	"github.com/floyoops/flo-go/backend/internal/ui/http/home"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
)

type ServerFactory struct {
	rootPath         string
	httpAllowOrigins []string
	homeCtrl         home.HomeController
	contactCtrl      contact.ContactController
}

func NewServerFactory(
	rootPath string,
	HttpAllowOrigins []string,
	homeCtrl home.HomeController,
	contactCtrl contact.ContactController,
) *ServerFactory {
	return &ServerFactory{
		rootPath:         rootPath,
		httpAllowOrigins: HttpAllowOrigins,
		homeCtrl:         homeCtrl,
		contactCtrl:      contactCtrl,
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

	router(e, f.homeCtrl, f.contactCtrl)

	return e
}
