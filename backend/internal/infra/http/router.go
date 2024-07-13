package http

import (
	"github.com/floyoops/flo-go/backend/internal/infra"
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo, di *infra.Container) {
	e.GET("/", func(c echo.Context) error { return di.HomeController.GetHome(c) })
	e.GET("/contact", func(c echo.Context) error { return di.ContactController.GetContact(c) })
	e.POST("/post-a-new-message-contact", func(c echo.Context) error { return di.ContactController.PostContact(c) })
}
