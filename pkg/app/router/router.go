package router

import (
	"github.com/floyoops/flo-go/pkg/app/di"
	"github.com/labstack/echo/v4"
)

const (
	HOME    = "/"
	CONTACT = "/contact"
)

type Router struct {
	e  *echo.Echo
	di *di.Container
}

func NewRouter(e *echo.Echo, di *di.Container) *Router {
	return &Router{e: e, di: di}
}

func (r *Router) Build() {
	r.e.GET(HOME, func(c echo.Context) error { return r.di.HomeController.GetHome(c) })
	r.e.GET(CONTACT, func(c echo.Context) error { return r.di.ContactController.GetContact(c) })
	r.e.POST(CONTACT, func(c echo.Context) error { return r.di.ContactController.PostContact(c) })
}
