package router

import (
	"github.com/floyoops/flo-go/pkg/app"
	"github.com/floyoops/flo-go/pkg/app/controller"
	"github.com/labstack/echo/v4"
)

func Init(di *app.Container, e *echo.Echo) {
	setHomeController(di, e)
	setContactController(di, e)
}

func setHomeController(di *app.Container, e *echo.Echo) {
	e.GET(controller.HOME, func(c echo.Context) error { return di.HomeController.GetHome(c) })
}

func setContactController(di *app.Container, e *echo.Echo) {
	e.GET(controller.CONTACT, func(c echo.Context) error { return di.ContactController.GetContact(c) })
	e.POST(controller.CONTACT, func(c echo.Context) error { return di.ContactController.PostContact(c) })
}
