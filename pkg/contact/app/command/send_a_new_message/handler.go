package send_a_new_message

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Handle(command Command) bool {
	return true
}
