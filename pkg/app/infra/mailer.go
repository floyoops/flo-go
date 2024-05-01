package infra

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
)

type Mailer struct {
	Host string
	Port string
	User string
	Pass string
}

func NewMailer(Host, Port, User, Pass string) mailer.Mailer {
	return &Mailer{Host: Host, Port: Port, User: User, Pass: Pass}
}

func (m *Mailer) Send(email string, text string) (bool, error) {
	fmt.Printf("Message %s from email %s sended", email, text)
	return true, nil
}
