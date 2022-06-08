package cSystemLogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	// Add imports here
)

const (
	// Supported log levels
	DEBUG = "DEBUG"
	INFO  = "INFO"
)

type SystemLogger struct {
	application  string
	environment  string
	internalIP   string
	directOutput *os.File
	loggerPtr    *log.Logger
}

func New(application, environment, internalIP string, outputToFile bool) (systemLoggerPtr *SystemLogger) {
	//func New(application, environment, internalIP string, outputToFile, splitErrors bool) (systemLoggerPtr *SystemLogger) {

	systemLoggerPtr = &SystemLogger{
		application: application,
		environment: environment,
		internalIP:  internalIP,
		loggerPtr:   initialize(application, environment, internalIP, INFO, outputToFile),
	}

	systemLoggerPtr.loggerPtr.Println("System Logger has been initialized")

	return
}

func initialize(application, environment, internalIP, logLevel string, writeToFile bool) *log.Logger {

	var (
		directOutput io.Writer = os.Stdout
		err          error
	)

	if writeToFile {
		directOutput, err = os.OpenFile(getLogfileName(application, environment, internalIP), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
	}

	return createLogger(directOutput, application, environment, internalIP, logLevel)
}

func getLogfileName(application, environment, internalIP string) string {
	return fmt.Sprintf("%v_%v_%v_%v.log", application, environment, internalIP, time.Now().Format("2006-01-02_15:04:05.12345_-0700"))
}

func createLogger(directOutput io.Writer, application, environment, internalIP, logLevel string) *log.Logger {
	return log.New(directOutput, fmt.Sprintf("%v.%v.%v.%v:", application, environment, internalIP, logLevel), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)
}

func (sl SystemLogger) TurnDebugOn() {
	createLogger(sl.directOutput, sl.application, sl.environment, sl.internalIP, DEBUG)
	sl.loggerPtr.Println("System Logger in Debug mode")
}

func (sl SystemLogger) TurnDebugOff() {
	createLogger(sl.directOutput, sl.application, sl.environment, sl.internalIP, INFO)
	sl.loggerPtr.Println("System Logger in Info mode")
}
