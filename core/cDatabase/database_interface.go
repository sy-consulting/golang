package cDatabase

type DatabaseIF interface {
	//initLogger(io.Writer, string) *SystemLogger
	TurnDebugOn()
}

func TurnDebugOn(message DatabaseIF) {
	message.TurnDebugOn()
}
