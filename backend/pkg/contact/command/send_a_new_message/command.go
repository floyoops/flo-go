package send_a_new_message

import "github.com/floyoops/flo-go/backend/pkg/contact/domain/model"

type Command struct {
	Name    string
	Email   *model.Email
	Message string
}
