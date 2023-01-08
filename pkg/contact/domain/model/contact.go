package model

type Contact struct {
	Name    string
	Email   string
	Message string
}

func NewContact(name string, email string, message string) *Contact {
	return &Contact{Name: name, Email: email, Message: message}
}
