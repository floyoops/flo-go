package send_a_new_message

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/event"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/floyoops/flo-go/backend/pkg/core"
)

type SendANewMessageCommandHandler struct {
	contactRepository repository.ContactRepository
	mailer            mailer.Mailer
	contactEmailApp   *model.Email
}

func NewHandler(contactRepository repository.ContactRepository, mailer mailer.Mailer, contactFromEmail *model.Email) *SendANewMessageCommandHandler {
	return &SendANewMessageCommandHandler{contactRepository, mailer, contactFromEmail}
}

func (h SendANewMessageCommandHandler) Handle(command bus.Command) ([]bus.Event, error) {
	cmd, ok := command.(*SendANewMessageCommand)
	if !ok {
		return nil, errors.New("invalid command type for SendANewMessageCommandHandler")
	}
	contact := model.NewContact(
		core.NewIdentifier(),
		cmd.Name,
		cmd.Email,
		cmd.Message,
	)

	errRepo := h.contactRepository.Create(contact)
	if errRepo != nil {
		return nil, errRepo
	}

	to := model.NewEmailList([]*model.Email{h.contactEmailApp})
	subject := "App flo-go new message from " + contact.Name
	body := fmt.Sprintf("name: %s\nemail: %s\nmessage: %s", contact.Name, contact.Email, contact.Message)
	err := h.mailer.Send(h.contactEmailApp, to, subject, body)
	if err != nil {
		return nil, err
	}

	return []bus.Event{event.NewANewMessageHasBeenSendEvent(cmd.Name, cmd.Email, cmd.Message)}, nil
}
