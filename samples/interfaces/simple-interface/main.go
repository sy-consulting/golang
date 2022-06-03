package main

import (
	"fmt"
	// Add imports here
)

const (
	LOGMESSAGEPREFIX = "Golang/samples/interfaces"
	// Add Constants here
)

type sMessage interface {
	natsCall()
}

func natsCall(message sMessage) {
	message.natsCall()
}

type realSMessage struct {
	owner string
}

func (realSMessage) natsCall() {
	fmt.Println("NATS is responding.")
}

type FakeSMessage struct {
	owner string
}

func (FakeSMessage) natsCall() {
	fmt.Println("I'm an imposter.")
}

func main() {

	realCalls := []realSMessage{{"Scott"}}
	for _, call := range realCalls {
		call.natsCall()
	}

	fakeCalls := []FakeSMessage{{"Scott"}}
	for _, call := range fakeCalls {
		call.natsCall()
	}

	calls := []sMessage{realSMessage{"Scott"}, FakeSMessage{"Scott"}}
	for _, call := range calls {
		natsCall(call)
	}
}
