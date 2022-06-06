/*
messageManager oversees all the shared services, such as logging, errors, messaging, etc.
	The start-up parameters can be supplied from a file or on the command line. The file contents is in JSON format and
	must be in the same directory as the binary using the messageManager services.
		Example: Business Service X (Path: ../BusinessService/X) is using messageManager logging, so the init file must be in the
		../BusinessService directory.

	All errors are in JSON format.
	JSON Fields:
		application      Blah Blah

Notes:
	None
*/
package core

import (
	"fmt"
	core "github.com/sy-consulting/golang/core/cSystemInfo"
	"github.com/sy-consulting/golang/samples/interfaces/interface-manager/messageManager"
)

const (
	CORE_MESSAGES = "MESSAGES"
	CORE_DATABASE = "DATABASE"
)

type Manager struct {
	application string
	environment string
	internalIP  string
}

// Core.New
func New(application, environment string) (coreManagerPtr *Manager) {

	coreManagerPtr = &Manager{
		application: application,
		environment: environment,
	}

	fmt.Println("Example one - single and slice calls")
	// Call to Mock New
	core.SystemInfoMock{}.GetIP()
	//Multiple calls to Mock New
	mockCalls := []core.SystemInfoMock{{}, {}}
	for _, call := range mockCalls {
		call.GetIP()
		fmt.Printf("IP: %v\n", call.InternalIP)
	}
	fmt.Println("Example one - done")

	fmt.Println("Example two - single and slice calls")
	// Call to Mock New
	coreManagerPtr.internalIP = core.SystemInfo{}.GetIP()
	fmt.Printf("IP: %v\n", coreManagerPtr.internalIP)

	//Multiple calls to Mock New
	Calls := []core.SystemInfo{{}, {}}
	for _, call := range Calls {
		coreManagerPtr.internalIP = call.GetIP()
		fmt.Printf("IP: %v\n", coreManagerPtr.internalIP)
	}
	fmt.Println("Example one - done")

	fmt.Println("===================")
	fmt.Println("Using Interface")
	fmt.Println("===================")

	fmt.Println("Example running both")
	// Call to Real New
	cSM :=
	cSM.New()
	// Call to Mock New
	cSM =
	// Multiple calls to Real New
	calls := []messageManager.SMessage{messageManager.RealSMessage{"Scott"}, messageManager.MockSMessage{"Scott"}}
	for _, call := range calls {
		call.New()
		call.NatsCall()
	}
	fmt.Println("Example running both - done\n")


	return
}
