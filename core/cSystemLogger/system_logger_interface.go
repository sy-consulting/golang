package cSystemLogger

type SystemLoggerIf interface {
	TurnDebugOn()
}

func TurnDebugOn(message SystemLoggerIf) {
	message.TurnDebugOn()
}
