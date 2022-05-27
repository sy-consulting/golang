package core

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type FakeSMessage struct {
	owner string
}

func (FakeSMessage) natsCall() {
	fmt.Println("I'm an imposter.")
}
