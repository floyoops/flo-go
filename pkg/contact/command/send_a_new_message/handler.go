package send_a_new_message

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/pkg/contact/repository"
	"github.com/floyoops/flo-go/pkg/core"
)

type Handler struct {
	contactRepository repository.ContactRepository
	mailer            mailer.Mailer
	contactEmailApp   *model.Email
}

func NewHandler(contactRepository repository.ContactRepository, mailer mailer.Mailer, contactFromEmail *model.Email) *Handler {
	return &Handler{contactRepository, mailer, contactFromEmail}
}

func (h *Handler) Handle(command Command) (bool, error) {
	contact := model.NewContact(
		core.NewIdentifier(),
		command.Name,
		command.Email,
		command.Message,
	)

	err2 := h.contactRepository.Create(contact)
	if err2 != nil {
		return false, err2
	}

	to := model.NewEmailList([]*model.Email{h.contactEmailApp})
	subject := "App flo-go new message from " + contact.Name
	body := fmt.Sprintf("name: %s\nemail: %s\nmessage: %s", contact.Name, contact.Email, contact.Message)
	_, err := h.mailer.Send(h.contactEmailApp, to, subject, body)
	if err != nil {
		return false, err
	}

	return true, nil
}
