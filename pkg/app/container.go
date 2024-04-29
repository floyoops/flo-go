package app

import (
	"github.com/floyoops/flo-go/pkg/app/controller"
	"github.com/floyoops/flo-go/pkg/contact/app/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
)

type Container struct {
	mailer                       mailer.Mailer
	SendNewMessageCommandHandler *send_a_new_message.Handler
	HomeController               controller.HomeController
	ContactController            controller.ContactController
}

func NewContainer() *Container {
	newMailer := mailer.NewMailer()
	sendANewsMessageCommandHandler := send_a_new_message.NewHandler(newMailer)
	homeController := controller.NewHomeController()
	contactController := controller.NewContactController(sendANewsMessageCommandHandler)
	return &Container{
		mailer:                       newMailer,
		SendNewMessageCommandHandler: sendANewsMessageCommandHandler,
		ContactController:            contactController,
		HomeController:               homeController,
	}
}
