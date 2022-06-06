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

	fmt.Println("Example one - single and slice calls")
	// Call to Mock New
	core.MockSMessage{}.New()
	// Multiple calls to Mock New
	mockCalls := []core.MockSMessage{{"SCOTT"}}
	for _, call := range mockCalls {
		call.New()
		call.NatsCall()
	}
	fmt.Println("Example one - done\n")

	fmt.Println("Example two - single and slice calls")
	// Call to Real New
	core.RealSMessage{}.New()
	// Multiple calls to Real New
	realCalls := []core.RealSMessage{{"SCOTT"}}
	for _, call := range realCalls {
		call.New()
		call.NatsCall()
	}
	fmt.Println("Example two - done\n")

	fmt.Println("Example running both")
	// Call to Real New
	cSM := core.SMessage(core.RealSMessage{})
	cSM.New()
	// Call to Mock New
	cSM = core.SMessage(core.MockSMessage{})
	cSM.New()
	// Multiple calls to Real New
	calls := []core.SMessage{core.RealSMessage{"Scott"}, core.MockSMessage{"Scott"}}
	for _, call := range calls {
		call.New()
		call.NatsCall()
	}
	fmt.Println("Example running both - done\n")

	fmt.Println("Example running only Real")
	// Call to Real New
	cSM = core.SMessage(core.RealSMessage{})
	cSM.New()
	calls = []core.SMessage{core.RealSMessage{"Scott"}}
	for _, call := range calls {
		call.New()
		call.NatsCall()
	}
	fmt.Println("Example running only Real - done\n")

	fmt.Println("Example running only mock")
	// Call to Mock New
	cSM = core.SMessage(core.MockSMessage{})
	cSM.New()
	calls = []core.SMessage{core.MockSMessage{"Scott"}}
	for _, call := range calls {
		call.New()
		call.NatsCall()
	}
	fmt.Println("Example running only mock - done")

	fmt.Println("Test has completed!")

}
