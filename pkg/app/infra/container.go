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
	config := NewConfig(rootPath)
	databaseDns := config.DatabaseUser + ":" + config.DatabasePassword + "@tcp(" + config.DatabaseHost + ":" + config.DatabasePort + ")/" + config.DatabaseName
	database := NewDatabase(databaseDns)
	contactRepository := NewContactMysqlRepository(database)
	newMailer := NewMailer(config.SmtpHost, config.SmtpPort, config.SmtpUsername, config.SmtpPassword)
	sendANewsMessageCommandHandler := send_a_new_message.NewHandler(contactRepository, newMailer, config.ContactEmailApp)
	homeController := home.NewHomeController()
	contactController := contact.NewContactController(sendANewsMessageCommandHandler)
	return &Container{
		Config:                       config,
		mailer:                       newMailer,
		SendNewMessageCommandHandler: sendANewsMessageCommandHandler,
		ContactController:            contactController,
		HomeController:               homeController,
	}
}
