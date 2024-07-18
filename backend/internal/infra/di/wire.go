//go:build wireinject
// +build wireinject

package di

import (
	"github.com/floyoops/flo-go/backend/config"
	"github.com/floyoops/flo-go/backend/internal"
	"github.com/floyoops/flo-go/backend/internal/infra/http"
	"github.com/floyoops/flo-go/backend/internal/ui/http/contact"
	"github.com/floyoops/flo-go/backend/internal/ui/http/home"
	"github.com/floyoops/flo-go/backend/pkg/contact/command/send_a_new_message"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/contact/infra"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/floyoops/flo-go/backend/pkg/database"
	"github.com/google/wire"
)

func provideServerFactory(
	config *config.Config,
	homeCtrl home.HomeController,
	contactCtrl contact.ContactController) *http.ServerFactory {
	return http.NewServerFactory(config.RootPath, config.HttpAllowOrigins, homeCtrl, contactCtrl)
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

var (
	databaseWiring = wire.NewSet(
		infra.NewContactMysqlRepository,
		wire.Bind(new(repository.ContactRepository), new(*infra.ContactMysqlRepository)),
	)
)

func BuildApp() (*internal.App, error) {

	wire.Build(
		config.NewConfig,
		provideDatabase,
		provideMailer,
		databaseWiring,
		provideContactFromEmail,
		send_a_new_message.NewHandler,
		home.NewHomeController,
		contact.NewContactController,
		provideServerFactory,
		internal.NewApp,
	)
	return &internal.App{}, nil
}
