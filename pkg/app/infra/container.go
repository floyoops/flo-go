package infra

import (
	"github.com/floyoops/flo-go/pkg/app/ui/http/contact"
	"github.com/floyoops/flo-go/pkg/app/ui/http/home"
	"github.com/floyoops/flo-go/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
)

type Container struct {
	Config                       *Config
	mailer                       mailer.Mailer
	SendNewMessageCommandHandler *send_a_new_message.Handler
	HomeController               home.HomeController
	ContactController            contact.ContactController
}

func NewContainer(rootPath string) *Container {
	newConfig := NewConfig(rootPath)
	newMailer := NewMailer(newConfig.SmtpHost, newConfig.SmtpPort, newConfig.SmtpUsername, newConfig.SmtpPassword)
	sendANewsMessageCommandHandler := send_a_new_message.NewHandler(newMailer)
	homeController := home.NewHomeController()
	contactController := contact.NewContactController(sendANewsMessageCommandHandler)
	return &Container{
		Config:                       newConfig,
		mailer:                       newMailer,
		SendNewMessageCommandHandler: sendANewsMessageCommandHandler,
		ContactController:            contactController,
		HomeController:               homeController,
	}
}
