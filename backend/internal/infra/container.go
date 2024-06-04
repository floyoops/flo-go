package infra

import (
	"github.com/floyoops/flo-go/config"
	"github.com/floyoops/flo-go/internal/ui/http/contact"
	"github.com/floyoops/flo-go/internal/ui/http/home"
	"github.com/floyoops/flo-go/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/pkg/contact/infra"
	database2 "github.com/floyoops/flo-go/pkg/database"
	mailer2 "github.com/floyoops/flo-go/pkg/mailer"
)

type Container struct {
	Config                       *config.Config
	mailer                       mailer.Mailer
	SendNewMessageCommandHandler *send_a_new_message.Handler
	HomeController               home.HomeController
	ContactController            contact.ContactController
}

func NewContainer(rootPath string) *Container {
	config := config.NewConfig(rootPath)
	databaseDns := config.DatabaseUser + ":" + config.DatabasePassword + "@tcp(" + config.DatabaseHost + ":" + config.DatabasePort + ")/" + config.DatabaseName
	database := database2.NewDatabase(databaseDns)
	contactRepository := infra.NewContactMysqlRepository(database)
	newMailer := mailer2.NewMailer(config.SmtpHost, config.SmtpPort, config.SmtpUsername, config.SmtpPassword)
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
