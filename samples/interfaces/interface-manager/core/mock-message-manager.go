package core

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type MockSMessage struct {
	Application string
}

func (MockSMessage) New() {
	fmt.Println("I'm an imposter. My Name is Mr. Mock New.")
}

func (MockSMessage) NatsCall() {
	fmt.Println("I'm an imposter. My Name is Mr. Mock.")
}
