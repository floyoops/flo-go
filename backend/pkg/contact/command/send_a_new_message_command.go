package command

import (
	"errors"
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/event"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/floyoops/flo-go/backend/pkg/core"
)

type SendANewMessageCommand struct {
	Name    string
	email   string
	Message string
}

func NewSendANewMessageCommand(name string, email *model.Email, message string) *SendANewMessageCommand {
	return &SendANewMessageCommand{Name: name, email: email.String(), Message: message}
}

func (c *SendANewMessageCommand) Identifier() bus.CommandIdentifier {
	return bus.NewIdentifierFromCommand(c)
}

func (c *SendANewMessageCommand) Email() (*model.Email, error) {
	return model.NewEmail(c.email)
}

type SendANewMessageCommandHandler struct {
	contactRepository repository.ContactRepository
}

func NewSendANewMessageCommandHandler(contactRepository repository.ContactRepository) *SendANewMessageCommandHandler {
	return &SendANewMessageCommandHandler{contactRepository}
}

func (h SendANewMessageCommandHandler) Handle(command bus.Command) ([]bus.Event, error) {
	cmd, ok := command.(*SendANewMessageCommand)
	if !ok {
		return nil, errors.New("invalid command type for SendANewMessageCommandHandler")
	}
	cmdEmail, err := cmd.Email()
	if err != nil {
		return nil, err
	}

	contact := model.NewContact(
		core.NewIdentifier(),
		cmd.Name,
		cmdEmail,
		cmd.Message,
	)

	errRepo := h.contactRepository.Create(contact)
	if errRepo != nil {
		return nil, errRepo
	}

	return []bus.Event{event.NewANewMessageHasBeenSendEvent(cmd.Name, cmdEmail, cmd.Message)}, nil
}
