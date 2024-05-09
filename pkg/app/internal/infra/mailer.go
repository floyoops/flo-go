package infra

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
	"github.com/floyoops/flo-go/pkg/contact/domain/model"
	"net/smtp"
	"strconv"
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

func (m *Mailer) Send(from *model.Email, to *model.EmailList, subject string, body string) (bool, error) {

	auth := smtp.PlainAuth("", m.User, m.Pass, m.Host)
	msg := []byte("To: " + to.ToArray()[0].String() + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	port, _ := strconv.Atoi(m.Port)
	err := smtp.SendMail(m.Host+":"+fmt.Sprintf("%d", port), auth, from.String(), to.ToArrayString(), msg)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de l'e-mail:", err)
		return false, err
	}

	fmt.Printf("Email from %s sended", from)
	return true, nil
}
