//go:build wireinject
// +build wireinject

package di

import (
	"github.com/floyoops/flo-go/backend/config"
	"github.com/floyoops/flo-go/backend/internal"
	"github.com/floyoops/flo-go/backend/internal/infra/http"
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact"
	"github.com/floyoops/flo-go/backend/internal/ui/http/home"
	"github.com/floyoops/flo-go/backend/pkg/bus"
	"github.com/floyoops/flo-go/backend/pkg/bus/middleware"
	"github.com/floyoops/flo-go/backend/pkg/contact/command"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/event"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/contact/infra"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/floyoops/flo-go/backend/pkg/database"
	"github.com/floyoops/flo-go/backend/pkg/logger"
	"github.com/google/wire"
)

func provideServerFactory(
	config *config.Config,
	routes []http.Route,
) *http.ServerFactory {
	return http.NewServerFactory(config.RootPath, config.HttpAllowOrigins, routes)
}

func provideDatabase(config *config.Config) *database.Database {
	return database.NewDatabase(config.GetDatabaseDns())
}

func provideMailer(config *config.Config) mailer.Mailer {
	return infra.NewMailer(config.SmtpHost, config.SmtpPort, config.SmtpUsername, config.SmtpPassword)
}

func provideContactFromEmail(config *config.Config) *model.Email {
	return config.ContactEmailApp
}

func provideApp(serverFactory *http.ServerFactory, logger logger.Logger, config *config.Config) *internal.App {
	app, err := internal.NewApp(serverFactory, logger, config.ServerPortHttp)
	if err != nil {
		panic(err)
	}
	return app
}

func provideCommandBus(
	SendANewMessageCommandHandler *command.SendANewMessageCommandHandler,
	ANewMessageHasBeenSendEventHandler *event.ANewMessageHasBeenSendEventHandler,
) *bus.CommandBus {
	eventBus := bus.NewSimpleEventBus()
	eventBus.RegisterHandler(&event.ANewMessageHasBeenSendEvent{}, ANewMessageHasBeenSendEventHandler)

	commandBus := bus.NewCommandBus(eventBus)
	commandBus.Use(middleware.LoggingMiddleware(logger.NewZapLogger()))
	commandBus.RegisterHandler(&command.SendANewMessageCommand{}, SendANewMessageCommandHandler)
	return commandBus
}

var (
	databaseWiring = wire.NewSet(
		infra.NewContactMysqlRepository,
		wire.Bind(new(repository.ContactRepository), new(*infra.ContactMysqlRepository)),
	)
	loggerWiring = wire.NewSet(
		logger.NewZapLogger,
		wire.Bind(new(logger.Logger), new(*logger.ZapLogger)),
	)
)

func BuildApp() (*internal.App, error) {

	wire.Build(
		config.NewConfig,
		loggerWiring,
		databaseWiring,
		provideDatabase,
		provideMailer,
		provideContactFromEmail,
		command.NewSendANewMessageCommandHandler,
		event.NewMessageHasBeenSendEventHandler,
		provideCommandBus,
		home.NewHomeController,
		contact.NewContactController,
		http.NewRoutes,
		provideServerFactory,
		provideApp,
	)
	return &internal.App{}, nil
}
