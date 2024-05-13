package contact

import (
	"errors"
	"github.com/floyoops/flo-go/pkg/app/internal/ui/http/contact/dto"
	"github.com/floyoops/flo-go/pkg/app/internal/ui/http/contact/view"
	"github.com/floyoops/flo-go/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/pkg/contact/repository"
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
	sendNewMessageCommandHandler *send_a_new_message.Handler
}

func NewContactController(sendNewMessageCommandHandler *send_a_new_message.Handler) ContactController {
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

	if errorsView := contactDto.Validate(); errorsView != nil {
		dataView := view.NewContactView(contactDto, &errorsView, false)
		return c.Render(http.StatusBadRequest, contactPage, dataView)
	}

	email, err := model.NewEmail(contactDto.Email)
	if err != nil {
		log.Errorf(err.Error())
	}

	err = ctl.sendNewMessageCommandHandler.Handle(send_a_new_message.Command{
		Name:    contactDto.Name,
		Email:   email,
		Message: contactDto.Message,
	})

	if err != nil {
		log.Errorf(err.Error())
		errorsView := make(map[string]string)
		if errors.Is(err, repository.ErrOnSaveContact) {
			errorsView = map[string]string{"error": "une erreur est survenue pendant la sauvegarde veuillez réessayer ultérieurement"}
		} else if errors.Is(err, mailer.ErrOnSend) {
			errorsView = map[string]string{"error": "une erreur est survenue pendant l envoie du mail veuillez réessayer ultérieurement"}
		} else {
			errorsView = map[string]string{"error": "une erreur est survenue veuillez réessayer ultérieurement"}
		}
		dataView := view.NewContactView(contactDto, &errorsView, false)
		return c.Render(http.StatusInternalServerError, contactPage, dataView)
	}

	dataView := view.NewContactView(contactDto, &map[string]string{}, true)
	return c.Render(http.StatusCreated, contactPage, dataView)
}
