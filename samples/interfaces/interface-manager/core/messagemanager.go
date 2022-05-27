package core

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type realSMessage struct {
	owner string
}

func (realSMessage) natsCall() {
	fmt.Println("NATS is responding.")
}
