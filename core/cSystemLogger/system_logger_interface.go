package cSystemLogger

type SystemLoggerIF interface {
	//initLogger(io.Writer, string) *SystemLogger
	TurnDebugOn()
}

//func initLogger(slIF SystemLoggerIF, infoHandle io.Writer, msgType string) (mySystemLogger *SystemLogger) {
//	return slIF.initLogger(infoHandle, msgType)
//}

func TurnDebugOn(message SystemLoggerIF) {
	message.TurnDebugOn()
}
