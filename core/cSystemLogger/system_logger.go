package cSystemLogger

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type SystemLogger struct {
	Application string
}

func New() {
	fmt.Println("I'm the real McCoy. My Name is Mr. New Second.")
}

func (SystemLogger) TurnDebugOn() {
	fmt.Println("I'm the real McCoy. My Name is Mr. Second.")
}
