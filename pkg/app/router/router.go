package router

import (
	"github.com/floyoops/flo-go/pkg/app/controller"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	setHomeController(e)
	setContactController(e)
}

func setHomeController(e *echo.Echo) {
	home := controller.NewHomeController()
	e.GET(controller.HOME, func(c echo.Context) error { return home.GetHome(c) })
}

func setContactController(e *echo.Echo) {
	contact := controller.NewContactController()
	e.GET(controller.CONTACT, func(c echo.Context) error { return contact.GetContact(c) })
	e.POST(controller.CONTACT, func(c echo.Context) error { return contact.PostContact(c) })
}
