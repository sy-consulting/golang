/*
System Info retrieves information in the SYC format

RESTRICTIONS:
    None

NOTES:
    None
*/
package cSystemLogger

import (
	"fmt"
	"io"
	"log"
)

type SystemLoggerMock struct {
	application string
	environment string
	internalIP  string
	lLevel      string
	logger      *log.Logger
	lHandler    io.Writer
}

// SystemLoggerMock.TurnDebugOn
func (SystemLoggerMock) TurnDebugOn() {
	fmt.Println("Debug logging is on - Gotch, I'm Mr. Mock!")
}
