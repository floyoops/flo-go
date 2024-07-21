package internal

import (
	"errors"
	"fmt"
	"github.com/floyoops/flo-go/backend/internal/infra/http"
	"github.com/floyoops/flo-go/backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo     *echo.Echo
	logger   logger.Logger
	portHttp int
}

func NewApp(serverFactory *http.ServerFactory, logger logger.Logger, portHttp int) (*App, error) {
	return &App{echo: serverFactory.Build(), logger: logger, portHttp: portHttp}, nil
}

func (a *App) Start() error {
	a.logger.Infof("Starting server on port %d", a.portHttp)
	if err := a.echo.Start(fmt.Sprintf(":%d", a.portHttp)); err != nil {
		a.logger.Errorf("Error on start server on port %d, %s", a.portHttp, err.Error())
		return errors.New("Error on start server " + err.Error())
	}
	a.logger.Infof("Server started on port %d", a.portHttp)
	return nil
}
