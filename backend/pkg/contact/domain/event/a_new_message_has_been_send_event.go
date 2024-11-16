package event

import (
	"errors"
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
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

type ANewMessageHasBeenSendEventHandler struct {
	mailer          mailer.Mailer
	contactEmailApp *model.Email
}

func NewMessageHasBeenSendEventHandler(mailer mailer.Mailer, contactFromEmail *model.Email) *ANewMessageHasBeenSendEventHandler {
	return &ANewMessageHasBeenSendEventHandler{mailer, contactFromEmail}
}

func (h ANewMessageHasBeenSendEventHandler) Handle(event bus.Event) error {
	evt, ok := event.(*ANewMessageHasBeenSendEvent)
	if !ok {
		return errors.New("invalid event type for ANewMessageHasBeenSendEventHandler")
	}

	to := model.NewEmailList([]*model.Email{h.contactEmailApp})
	subject := "App flo-go new message from " + evt.Name
	body := evt.Message
	err := h.mailer.Send(h.contactEmailApp, to, subject, body)
	if err != nil {
		return err
	}

	return nil
}
