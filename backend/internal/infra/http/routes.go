package http

import (
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact"
	"github.com/floyoops/flo-go/backend/internal/ui/http/home"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Route struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

func NewRoutes(hc home.HomeController, cc contact.ContactController) []Route {
	return []Route{
		{http.MethodGet, "/", hc.GetHome},
		{http.MethodGet, "/contact", cc.GetContact},
		{http.MethodPost, "/post-a-new-message-contact", cc.PostContact},
	}
}
