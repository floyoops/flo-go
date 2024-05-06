package infra

import (
	"fmt"
	"github.com/floyoops/flo-go/pkg/contact/domain/mailer"
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

func (m *Mailer) Send(from string, to []string, subject string, body string) (bool, error) {

	auth := smtp.PlainAuth("", m.User, m.Pass, m.Host)
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	port, _ := strconv.Atoi(m.Port)
	err := smtp.SendMail(m.Host+":"+fmt.Sprintf("%d", port), auth, from, to, msg)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de l'e-mail:", err)
		return false, err
	}

	fmt.Printf("Email from %s sended", from)
	return true, nil
}
