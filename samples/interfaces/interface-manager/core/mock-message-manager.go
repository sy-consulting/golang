package core

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type MockSMessage struct {
	Owner string
}

func (MockSMessage) NatsCall() {
	fmt.Println("I'm an imposter. My Name is Mr. Mock.")
}
