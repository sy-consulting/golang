package cDatabase

import (
	"fmt"
)

const (
// Add Constants here
)

type DatabaseMock struct {
	Application string
}

func (DatabaseMock) TurnDebugOn() {
	fmt.Println("I'm an imposter. My Name is Mr. Database Mock.")
}
