package send_a_new_message

import "github.com/floyoops/flo-go/backend/pkg/contact/domain/model"

type SendANewMessageCommand struct {
	Name    string
	Email   *model.Email
	Message string
}

func NewSendANewMessageCommand(name string, email *model.Email, message string) *SendANewMessageCommand {
	return &SendANewMessageCommand{Name: name, Email: email, Message: message}
}
