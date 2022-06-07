/*
General description of the purpose of the go file.

RESTRICTIONS:
    AWS functions:
    * Program must have access to a .aws/credentials file in the default location.
    * This will only access system parameters that start with '/sote' (ROOTPATH).
    * {Enter other restrictions here for AWS

    {Other catagories of restrictions}
    * {List of restrictions for the catagory

NOTES:
    {Enter any additional notes that you believe will help the next developer.}
*/
package cSystemLogger

import (
	"fmt"
	"io"
	"log"
)

const (
	// Supported log levels
	DEBUG_LEVEL = "DEBUG"
	INFO_LEVEL  = "INFO"
)

type SystemLogger struct {
	application string
	environment string
	internalIP  string
	lLevel      string
	logger      *log.Logger
	lHandler    io.Writer
}

func New(application, environment, internalIP string) (systemLoggerPtr *SystemLogger) {
	//func New(application, environment, internalIP, ) (systemLoggerPtr *SystemLogger) {

	systemLoggerPtr = &SystemLogger{
		application: application,
		environment: environment,
		internalIP:  internalIP,
		lLevel:      INFO_LEVEL,
		logger:      nil,
	}

	systemLoggerPtr.logger = log.New(systemLoggerPtr.lHandler, fmt.Sprintf("%v.%v.%v:", systemLoggerPtr.application, systemLoggerPtr.environment, systemLoggerPtr.internalIP), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC)

	return
}

// SystemLogger.TurnDebugOn
func (SystemLogger) TurnDebugOn() {
	fmt.Println("Debug logging is on - For Real")
}
