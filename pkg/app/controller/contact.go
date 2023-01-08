package controller

import (
	"github.com/floyoops/flo-go/pkg/app/dto"
	"github.com/floyoops/flo-go/pkg/app/view"
	"github.com/floyoops/flo-go/pkg/contact/app/command/send_a_new_message"
	"github.com/labstack/echo/v4"
	"net/http"
)

const contactPage string = "contact.page.html"

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
	return c.Render(http.StatusOK, contactPage, nil)
}

func (controller *contactController) PostContact(c echo.Context) error {
	contactDto := dto.NewContactDto()
	contactDto.Name = c.FormValue("name")
	contactDto.Email = c.FormValue("email")
	contactDto.Message = c.FormValue("message")

	if errors := contactDto.Validate(); errors != nil {
		dataView := view.NewContactView(contactDto, &errors, false)
		return c.Render(http.StatusBadRequest, contactPage, dataView)
	}

	result := send_a_new_message.NewHandler().Handle(send_a_new_message.Command{
		Name:    contactDto.Name,
		Email:   contactDto.Email,
		Message: contactDto.Message,
	})

	if result == false {
		errors := map[string]string{"error": "une erreur est survenue veuillez réessayer ultérieurement"}
		dataView := view.NewContactView(contactDto, &errors, false)
		return c.Render(http.StatusInternalServerError, contactPage, dataView)
	}

	dataView := view.NewContactView(contactDto, &map[string]string{}, true)
	return c.Render(http.StatusCreated, contactPage, dataView)
}
