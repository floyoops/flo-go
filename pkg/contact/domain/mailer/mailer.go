package mailer

type Mailer interface {
	Send(email string, text string) (bool, error)
}
