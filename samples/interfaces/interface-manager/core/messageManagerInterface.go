package core

const (
	LOGMESSAGEPREFIX = "Golang/samples/interfaces"
	// Add Constants here
)

type sMessage interface {
	natsCall()
}

func natsCall(message sMessage) {
	message.natsCall()
}
