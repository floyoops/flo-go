package mailer

import (
	"errors"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
)

type Mailer interface {
	Send(from *model.Email, to *model.EmailList, subject string, text string) error
}

var (
	ErrOnSend = errors.New("email on send")
)
