package contact

import (
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact/dto"
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact/view"
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/contact/command"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/core"
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
	commandBus *bus.CommandBus
}

func NewContactController(commandBus *bus.CommandBus) ContactController {
	return &contactController{commandBus: commandBus}
}

func (ctl *contactController) GetContact(c echo.Context) error {
	return c.Render(http.StatusOK, contactPage, map[string]interface{}{"NewId": core.NewIdentifier().String()})
}

func (ctl *contactController) PostContact(c echo.Context) error {
	contactDto := dto.NewContactDto()
	err := c.Bind(contactDto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if errorsView := contactDto.Validate(); errorsView != nil {
		dataView := view.NewContactView(contactDto, &errorsView, false)
		return c.JSON(http.StatusBadRequest, dataView)
	}

	email, err := model.NewEmail(contactDto.Email)
	if err != nil {
		log.Errorf(err.Error())
	}

	err = ctl.commandBus.Dispatch(command.NewSendANewMessageCommand(
		contactDto.Name,
		email,
		contactDto.Message,
	))

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
