package http

import (
	"errors"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	log.Errorf(err.Error())
	errorHttp := echo.NewHTTPError(http.StatusInternalServerError)
	if errors.Is(err, repository.ErrOnSaveContact) {
		errorHttp.Message = "une erreur est survenue pendant la sauvegarde veuillez réessayer ultérieurement"
	} else if errors.Is(err, mailer.ErrOnSend) {
		errorHttp.Message = "une erreur est survenue pendant l envoie du mail veuillez réessayer ultérieurement"
	} else {
		errorHttp.Message = "une erreur est survenue veuillez réessayer ultérieurement"
	}
	err = c.JSON(http.StatusInternalServerError, errorHttp)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
}
