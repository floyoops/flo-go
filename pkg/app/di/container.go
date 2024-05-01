package di

import (
	"github.com/floyoops/flo-go/pkg/app/config"
	"github.com/floyoops/flo-go/pkg/app/ui/http/contact"
	"github.com/floyoops/flo-go/pkg/app/ui/http/home"
	"github.com/floyoops/flo-go/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
)

type Container struct {
	Config                       *config.Config
	mailer                       mailer.Mailer
	SendNewMessageCommandHandler *send_a_new_message.Handler
	HomeController               home.HomeController
	ContactController            contact.ContactController
}

func NewContainer(rootPath string) *Container {
	newConfig := config.NewConfig(rootPath)
	newMailer := mailer.NewMailer()
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
