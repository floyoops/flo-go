package send_a_new_message

import (
	"errors"
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/event/a_new_message_has_been_send"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/floyoops/flo-go/backend/pkg/core"
)

type SendANewMessageCommandHandler struct {
	contactRepository repository.ContactRepository
}

func NewHandler(contactRepository repository.ContactRepository) *SendANewMessageCommandHandler {
	return &SendANewMessageCommandHandler{contactRepository}
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

	return []bus.Event{a_new_message_has_been_send.NewANewMessageHasBeenSendEvent(cmd.Name, cmd.Email, cmd.Message)}, nil
}
