/*
	This is SY-HOLDING's log wrapper for all logging in GOlang.

	This code is licensed under MIT Licence (https://mit-license.org/)
*/
package cSystemLogger

import (
	"fmt"
	"log"
	"os"
	// Add imports here
)

const (
	// Supported log levels
	DEBUG = "DEBUG"
	INFO  = "INFO"
	// Defined strings
	LOG_PREFIX        = "%v.%v.%v:"
	SETTING_LOG_LEVEL = "Setting System Logger to %v level"
	LOG_MESSAGE       = "%v %v"
)

type SystemLogger struct {
	application string
	environment string
	internalIP  string
	logLevel    string
	loggerPtr   *log.Logger
}

func New(application, environment, internalIP string) (systemLoggerPtr *SystemLogger) {

	systemLoggerPtr = &SystemLogger{
		application: application,
		environment: environment,
		internalIP:  internalIP,
		logLevel:    INFO,
		loggerPtr:   initialize(application, environment, internalIP, INFO),
	}

	return
}

func initialize(application, environment, internalIP, logLevel string) *log.Logger {

	return log.New(os.Stdout, fmt.Sprintf(LOG_PREFIX, application, environment, internalIP), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
}

func (sl SystemLogger) TurnDebugOn() {
	sl.logLevel = DEBUG
	sl.printLogLevel()
}

func (sl SystemLogger) TurnDebugOff() {
	sl.logLevel = INFO
	sl.printLogLevel()
}

func (sl SystemLogger) printLogLevel() {
	sl.loggerPtr.Printf(SETTING_LOG_LEVEL, sl.logLevel)
}

func (sl SystemLogger) LogMessage(message string) {
	sl.loggerPtr.Printf(LOG_MESSAGE, sl.logLevel, message)
}
