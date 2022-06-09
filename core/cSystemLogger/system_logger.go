/*
	This is SY-HOLDING's log wrapper for all logging in GOlang. If you think of a way to improve it, present the idea.

RESTRICTIONS:
	Do not call LogMethod in this file. It will cause an infinite loop.

LICENSE:
	This code is licensed under MIT Licence (https://mit-license.org/)
*/
package cSystemLogger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	// Add imports here
)

const (
	// Supported log levels
	INFO              = "INFO "
	INFO_VALUE  uint8 = 1
	DEBUG             = "DEBUG"
	DEBUG_VALUE uint8 = 2
	// Defined strings
	LOG_PREFIX        = "%v.%v.%v:"
	SETTING_LOG_LEVEL = "%v is the setting for System Logger"
	LOG_MESSAGE       = "%v %v"
)

type SystemLogger struct {
	application string
	environment string
	internalIP  string
	loggerPtr   *log.Logger
}

var (
	logLevel = INFO
	logValue = INFO_VALUE
)

func New(application, environment, internalIP string) (systemLoggerPtr *SystemLogger) {
	systemLoggerPtr = &SystemLogger{
		application: application,
		environment: environment,
		internalIP:  internalIP,
		loggerPtr:   initialize(application, environment, internalIP),
	}

	return
}

func initialize(application, environment, internalIP string) *log.Logger {

	return log.New(os.Stdout, fmt.Sprintf(LOG_PREFIX, application, environment, internalIP), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
}

func (sl SystemLogger) TurnDebugOn() {
	logLevel = DEBUG
	logValue = DEBUG_VALUE
	sl.printLogLevel()
}

func (sl SystemLogger) TurnDebugOff() {
	logLevel = INFO
	logValue = INFO_VALUE
	sl.printLogLevel()
}

func (sl SystemLogger) printLogLevel() {
	sl.loggerPtr.Printf(SETTING_LOG_LEVEL, logLevel)
}

func (sl SystemLogger) DebugMessage(message string) {
	if DEBUG_VALUE <= logValue {
		sl.LogMessage(message)
	}
}

func (sl SystemLogger) LogMessage(message string) {
	sl.loggerPtr.Printf(LOG_MESSAGE, logLevel, message)
}

func (sl SystemLogger) LogMethod(depthList ...int) {
	if DEBUG_VALUE <= logValue {
		var depth int
		if depthList == nil {
			depth = 1
		} else {
			depth = depthList[0]
		}
		function, _, _, _ := runtime.Caller(depth)
		sl.LogMessage(runtime.FuncForPC(function).Name())
	}
}
