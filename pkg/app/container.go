package app

import (
	"github.com/floyoops/flo-go/pkg/contact/app/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
)

type Container struct {
	mailer                       mailer.Mailer
	SendNewMessageCommandHandler *send_a_new_message.Handler
}

func NewContainer() *Container {
	newMailer := mailer.NewMailer()
	sendANewsMessageCommandHandler := send_a_new_message.NewHandler(newMailer)
	return &Container{
		mailer:                       newMailer,
		SendNewMessageCommandHandler: sendANewsMessageCommandHandler,
	}
}
