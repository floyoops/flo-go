package mailer

import "fmt"

type Mailer interface {
	Send(email string, text string) (bool, error)
}

type mailer struct {
}

func NewMailer() Mailer {
	return &mailer{}
}

func (m *mailer) Send(email string, text string) (bool, error) {
	fmt.Printf("Message %s from email %s sended", email, text)
	return true, nil
}
