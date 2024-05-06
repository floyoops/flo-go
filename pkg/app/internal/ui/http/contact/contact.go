package contact

import (
	"github.com/floyoops/flo-go/pkg/app/internal/ui/http/contact/dto"
	"github.com/floyoops/flo-go/pkg/app/internal/ui/http/contact/view"
	send_a_new_message2 "github.com/floyoops/flo-go/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/core"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

const contactPage string = "contact.page.html"

type ContactController interface {
	GetContact(c echo.Context) error
	PostContact(c echo.Context) error
}

type contactController struct {
	sendNewMessageCommandHandler *send_a_new_message2.Handler
}

func NewContactController(sendNewMessageCommandHandler *send_a_new_message2.Handler) ContactController {
	return &contactController{sendNewMessageCommandHandler: sendNewMessageCommandHandler}
}

func (ctl *contactController) GetContact(c echo.Context) error {
	return c.Render(http.StatusOK, contactPage, map[string]interface{}{"NewId": core.NewIdentifier().String()})
}

func (ctl *contactController) PostContact(c echo.Context) error {
	contactDto := dto.NewContactDto()
	contactDto.Name = c.FormValue("name")
	contactDto.Email = c.FormValue("email")
	contactDto.Message = c.FormValue("message")

	if errors := contactDto.Validate(); errors != nil {
		dataView := view.NewContactView(contactDto, &errors, false)
		return c.Render(http.StatusBadRequest, contactPage, dataView)
	}

	result, err := ctl.sendNewMessageCommandHandler.Handle(send_a_new_message2.Command{
		Name:    contactDto.Name,
		Email:   contactDto.Email,
		Message: contactDto.Message,
	})

	if err != nil {
		log.Errorf(err.Error())
	}

	if result == false {
		errors := map[string]string{"error": "une erreur est survenue veuillez réessayer ultérieurement"}
		dataView := view.NewContactView(contactDto, &errors, false)
		return c.Render(http.StatusInternalServerError, contactPage, dataView)
	}

	dataView := view.NewContactView(contactDto, &map[string]string{}, true)
	return c.Render(http.StatusCreated, contactPage, dataView)
}
