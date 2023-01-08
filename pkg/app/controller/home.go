package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HomeController interface {
	GetHome(c echo.Context) error
}

type homeController struct {
}

func NewHomeController() HomeController {
	return &homeController{}
}

func (controller *homeController) GetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"title": "HomePage!",
	})
}
