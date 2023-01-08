package send_a_new_message

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(command Command) bool {
	contact := model.NewContact(
		command.Name,
		command.Email,
		command.Message,
	)

	fmt.Printf("Message for %s handled", contact.Email)

	return true
}
