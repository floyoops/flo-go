package send_a_new_message

import (
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
)

type Handler struct {
	mailer mailer.Mailer
}

func NewHandler(mailer mailer.Mailer) *Handler {
	return &Handler{mailer}
}

func (h *Handler) Handle(command Command) bool {
	contact := model.NewContact(
		command.Name,
		command.Email,
		command.Message,
	)

	_, err := h.mailer.Send(contact.Email, contact.Message)
	if err != nil {
		return false
	}

	return true
}
