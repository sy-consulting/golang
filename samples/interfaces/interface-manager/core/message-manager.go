package core

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type RealSMessage struct {
	Owner string
}

func (RealSMessage) NatsCall() {
	fmt.Println("I'm the real McCoy. My Name is Mr. NATS.")
}
