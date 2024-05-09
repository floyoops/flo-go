package mailer

import "github.com/floyoops/flo-go/pkg/contact/domain/model"

type Mailer interface {
	Send(from *model.Email, to *model.EmailList, subject string, text string) (bool, error)
}
