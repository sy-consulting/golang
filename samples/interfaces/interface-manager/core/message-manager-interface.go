package core

const (
// Add Constants here
)

type SMessage interface {
	NatsCall()
}

func NatsCall(message SMessage) {
	message.NatsCall()
}
