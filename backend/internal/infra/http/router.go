package http

import (
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact"
	"github.com/floyoops/flo-go/backend/internal/ui/http/home"
	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo, homeCtrl home.HomeController, contactCtrl contact.ContactController) {
	e.GET("/", func(c echo.Context) error { return homeCtrl.GetHome(c) })
	e.GET("/contact", func(c echo.Context) error { return contactCtrl.GetContact(c) })
	e.POST("/post-a-new-message-contact", func(c echo.Context) error { return contactCtrl.PostContact(c) })
}
