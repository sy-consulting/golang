package messageManager

type SMessageIF interface {
	New()
	NatsCall()
}

func New(message SMessageIF) {
	message.New()
}

func NatsCall(message SMessageIF) {
	message.NatsCall()
}
