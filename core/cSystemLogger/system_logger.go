package cSystemLogger

import (
	"fmt"
	"log"
	"os"
	// Add imports here
)

const (
// Add Constants here
)

type SystemLogger struct {
	Logger *log.Logger
}

func New() (systemLoggerPtr *SystemLogger) {

	systemLoggerPtr = &SystemLogger{
		Logger: log.New(os.Stdout, fmt.Sprintf("%v.%v:", "TEST", "INFO"), log.Lmsgprefix|log.LstdFlags|log.Lmicroseconds|log.LUTC),
	}

	systemLoggerPtr.Logger.Println("System Logger has been initialized")
	//
	return
}

func (SystemLogger) TurnDebugOn() {
	fmt.Println("I'm the real McCoy. My Name is Mr. Second.")
}
