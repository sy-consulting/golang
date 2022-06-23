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
	// Testing
	TEST_LOGPANIC = "github.com/sy-consulting/golang/core/cSystemLogger.TestLogPanic"
	TEST_LOGFATAL = "github.com/sy-consulting/golang/core/cSystemLogger.TestLogFatal"
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

func (sl SystemLogger) IsDebugOn() (result bool) {
	if logLevel == DEBUG {
		return true
	}

	return false
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

func (sl SystemLogger) LogPanic(panicData []string, depthList ...int) {

	var (
		depth    int
		logPanic = log.New(os.Stderr, fmt.Sprintf(LOG_PREFIX, sl.application, sl.environment, sl.internalIP), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
	)
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, _, _, _ := runtime.Caller(depth)
	logPanic.Println(runtime.FuncForPC(function).Name())
	if runtime.FuncForPC(function).Name() != TEST_LOGPANIC {
		logPanic.Panic(panicData)
	}
}

func (sl SystemLogger) LogFatal(fatalData []string, depthList ...int) {

	var (
		depth    int
		logPanic = log.New(os.Stderr, fmt.Sprintf(LOG_PREFIX, sl.application, sl.environment, sl.internalIP), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
	)
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, _, _, _ := runtime.Caller(depth)
	logPanic.Println(runtime.FuncForPC(function).Name())
	if runtime.FuncForPC(function).Name() != TEST_LOGFATAL {
		logPanic.Fatalln(fatalData)
	}
}
