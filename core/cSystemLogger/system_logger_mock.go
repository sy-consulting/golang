package cSystemLogger

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type SystemLoggerMock struct {
	Application string
}

func (SystemLoggerMock) TurnDebugOn() {
	fmt.Println("I'm an imposter. My Name is Mr. Second Mock.")
}
