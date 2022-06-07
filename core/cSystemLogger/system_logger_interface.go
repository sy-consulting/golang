package cSystemLogger

type SystemLoggerIF interface {
	TurnDebugOn() string
}

func TurnDebugOn(slIF SystemLoggerIF) {
	slIF.TurnDebugOn()
}
