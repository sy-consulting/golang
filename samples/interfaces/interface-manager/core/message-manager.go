package core

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type RealSMessage struct {
	Application string
}

func (RealSMessage) New() {
	fmt.Println("I'm the real McCoy. My Name is Mr. New NATS.")
}

func (RealSMessage) NatsCall() {
	fmt.Println("I'm the real McCoy. My Name is Mr. NATS.")
}
