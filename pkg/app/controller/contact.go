package controller

import (
	"github.com/floyoops/flo-go/pkg/contact/app/command/send_a_new_message"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactController interface {
	GetContact(c echo.Context) error
	PostContact(c echo.Context) error
}

type contactController struct {
}

func NewContactController() ContactController {
	return &contactController{}
}

func (controller *contactController) GetContact(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func (controller *contactController) PostContact(c echo.Context) error {
	result := send_a_new_message.NewHandler().Handle(send_a_new_message.Command{
		Name:    "a",
		Email:   "b",
		Message: "c",
	})

	if result == false {
		return c.JSON(http.StatusInternalServerError, "error on send message")
	}
	return c.JSON(http.StatusOK, "ok")
}
