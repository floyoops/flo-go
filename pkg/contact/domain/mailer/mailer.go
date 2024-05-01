package mailer

type Mailer interface {
	Send(from string, to []string, subject string, text string) (bool, error)
}
