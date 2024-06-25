package infra

import (
	"github.com/floyoops/flo-go/backend/config"
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact"
	"github.com/floyoops/flo-go/backend/internal/ui/http/home"
	"github.com/floyoops/flo-go/backend/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/backend/pkg/contact/infra"
	database2 "github.com/floyoops/flo-go/backend/pkg/database"
	"github.com/floyoops/flo-go/backend/pkg/logger"
)

type Container struct {
	Config                       *config.Config
	Logger                       logger.Logger
	mailer                       mailer.Mailer
	SendNewMessageCommandHandler *send_a_new_message.Handler
	HomeController               home.HomeController
	ContactController            contact.ContactController
}

func NewContainer(rootPath string) *Container {
	configuration := config.NewConfig(rootPath)
	zapLogger := logger.NewZapLogger()
	databaseDns := configuration.DatabaseUser + ":" + configuration.DatabasePassword + "@tcp(" + configuration.DatabaseHost + ":" + configuration.DatabasePort + ")/" + configuration.DatabaseName
	database := database2.NewDatabase(databaseDns)
	contactRepository := infra.NewContactMysqlRepository(database)
	newMailer := infra.NewMailer(configuration.SmtpHost, configuration.SmtpPort, configuration.SmtpUsername, configuration.SmtpPassword)
	sendANewsMessageCommandHandler := send_a_new_message.NewHandler(contactRepository, newMailer, configuration.ContactEmailApp)
	homeController := home.NewHomeController()
	contactController := contact.NewContactController(sendANewsMessageCommandHandler)
	return &Container{
		Config:                       configuration,
		Logger:                       zapLogger,
		mailer:                       newMailer,
		SendNewMessageCommandHandler: sendANewsMessageCommandHandler,
		ContactController:            contactController,
		HomeController:               homeController,
	}
}
