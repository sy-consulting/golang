package cSystemLogger

type SystemLoggerIF interface {
	TurnDebugOn() string
}

func TurnDebugOn(siIF SystemLoggerIF) {
	siIF.TurnDebugOn()
}
