package infra

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	e  *echo.Echo
	di *Container
}

func NewRouter(e *echo.Echo, di *Container) *Router {
	return &Router{e: e, di: di}
}

func (r *Router) Build() {
	r.e.GET("/", func(c echo.Context) error { return r.di.HomeController.GetHome(c) })
	r.e.GET("/contact", func(c echo.Context) error { return r.di.ContactController.GetContact(c) })
	r.e.POST("/post-a-new-message-contact", func(c echo.Context) error { return r.di.ContactController.PostContact(c) })
}
