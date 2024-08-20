package telegram

var (
	enabled = false

	nc = Chat(NormalChat{})
	mc = Chat(MonitorChat{})
)

type Chat struct {
}

type NormalChat Chat
type MonitorChat Chat

func (nc NormalChat) Recipient() string {
	return chatID
}
func (mc MonitorChat) Recipient() string {
	return monitorChatID
}
