/*
core/logger/manager writes message to the local file system

The log message format is:
    {YYYY}/{MM}/{DD} {hh}:{mm}:{ssss}{+|-}{TZD}  {application}.{environment}.{hostname}.{MessageType}:{Message}
    example: 2020/06/16 22:26:42.1656-08:00 CLIENTMANAGER.STAGING.MYMACMINI.DEBUG:gitlab.com/nats.io/nats.go is being used

Notes:
	hostname must be unique for all computers
	Non-fatal log message go to stdout
	Fatal log message go to stderr
*/
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

const (
	DebugLevel    = "DEBUG"   // Support log levels
	InfoLevel     = "INFO"    // Support log levels
	PrefixMissing = "missing" // Support log levels
)

type LManager struct {
	application string
	environment string
	hostname    string
	lLevel      string
	logger      *log.Logger
	lHandler    io.Writer
	prefix      string
}

// Starts logging
func New(application, environment, logLevel string) (lManagerPtr *LManager) {
	// Initialize the values for Nats Manager
	lManagerPtr = &LManager{
		application: application,
		environment: environment,
		hostname:    "",
		lLevel:      logLevel,
		logger:      nil,
		lHandler:    nil,
		prefix:      "",
	}

	lManagerPtr.hostname, _ = os.Hostname()
	lManagerPtr.prefix = fmt.Sprintf("%v.%v.%v.%v:", lManagerPtr.application, lManagerPtr.environment, lManagerPtr.hostname)
	lManagerPtr.logger = log.New(lManagerPtr.lHandler, lManagerPtr.prefix, log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
}

func initLogger(infoHandle io.Writer, msgType string) {
	lMessage = log.New(infoHandle, fmt.Sprintf("%v.%v:", lPrefix, msgType), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
}

// This will publish a log message at the INFO level
func Info(tMessage string) {
	initLogger(os.Stdout, InfoLogLevel)
	logMessage.Println(tMessage)
}

// This will publish a log message at the DEBUG level
func Debug(tMessage string) {
	if logLevel == DebugLogLevel {
		initLogger(os.Stdout, DebugLogLevel)
		logMessage.Println(tMessage)
	}
}

// This will publish a log message at the DEBUG level for the function that is being executed.
func Method(depthList ...int) {
	if logLevel == DebugLogLevel {
		initLogger(os.Stdout, DebugLogLevel)
		var depth int
		if depthList == nil {
			depth = 1
		} else {
			depth = depthList[0]
		}
		function, _, _, _ := runtime.Caller(depth)
		functionName := runtime.FuncForPC(function).Name()
		logMessage.Println(functionName)
	}
}

// This will set the log level to DEBUG
func SetToDebug() {
	logLevel = DebugLogLevel
}

// This will set the log level to INFO
func SetToInfo() {
	logLevel = InfoLogLevel
}

// This will return the logging level (It doesn't follow return referring to the func declaration)
func GetLevel() string {
	return logLevel
}

// Allows a message prefix to be applied
// The prefix is forced to upper case
func SetMessagePrefix(prefix string) {
	MeessagePrefix = strings.ToUpper(prefix)
}
