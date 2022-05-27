package main

import (
	"fmt"
	"github.com/sy-consulting/golang/samples/interfaces/interface-manager/core"
)

const (
	LOGMESSAGEPREFIX = "Golang/samples/interfaces/interface-manager"
	// Add Constants here
)

// List type's here

var (
// Add Variables here for the file (Remember, they are global)
)

func main() {

	fmt.Println("Example one")
	mockCalls := []core.MockSMessage{{"SCOTT"}}
	for _, call := range mockCalls {
		call.NatsCall()
	}
	fmt.Println("Example one - done\n")

	fmt.Println("Example two")
	realCalls := []core.RealSMessage{{"SCOTT"}}
	for _, call := range realCalls {
		call.NatsCall()
	}
	fmt.Println("Example two - done\n")

	fmt.Println("Example running both")
	calls := []core.SMessage{core.RealSMessage{"Scott"}, core.MockSMessage{"Scott"}}
	for _, call := range calls {
		call.NatsCall()
	}
	fmt.Println("Example running both - done\n")

	fmt.Println("Example running only Real")
	calls = []core.SMessage{core.RealSMessage{"Scott"}}
	for _, call := range calls {
		call.NatsCall()
	}
	fmt.Println("Example running only Real - done\n")

	fmt.Println("Example running only mock")
	calls = []core.SMessage{core.MockSMessage{"Scott"}}
	for _, call := range calls {
		call.NatsCall()
	}
	fmt.Println("Example running only mock - done")

	fmt.Println("Test has completed!")

}
