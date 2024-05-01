package send_a_new_message

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
)

type Handler struct {
	mailer          mailer.Mailer
	contactEmailApp string
}

func NewHandler(mailer mailer.Mailer, contactFromEmail string) *Handler {
	return &Handler{mailer, contactFromEmail}
}

func (h *Handler) Handle(command Command) bool {
	contact := model.NewContact(
		command.Name,
		command.Email,
		command.Message,
	)

	to := []string{h.contactEmailApp}
	subject := "App flo-go new message from " + contact.Name
	body := fmt.Sprintf("name: %s\nemail: %s\nmessage: %s", contact.Name, contact.Email, contact.Message)
	_, err := h.mailer.Send(h.contactEmailApp, to, subject, body)
	if err != nil {
		return false
	}

	return true
}
