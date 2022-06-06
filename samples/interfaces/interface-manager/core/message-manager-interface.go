package core

const (
// Add Constants here
)

type SMessage interface {
	New()
	NatsCall()
}

func New(message SMessage) {
	message.New()
}

func NatsCall(message SMessage) {
	message.NatsCall()
}
