package a_new_message_has_been_send

import (
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
)

type ANewMessageHasBeenSendEvent struct {
	Name    string
	Email   string
	Message string
}

func NewANewMessageHasBeenSendEvent(name string, email *model.Email, message string) *ANewMessageHasBeenSendEvent {
	return &ANewMessageHasBeenSendEvent{Name: name, Email: email.String(), Message: message}
}

func (e *ANewMessageHasBeenSendEvent) Identifier() bus.EventIdentifier {
	return bus.NewIdentifierFromEvent(e)
}
