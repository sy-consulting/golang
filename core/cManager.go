/*
core oversees all the shared services, such as logging, errors, messaging, etc.
	The start-up parameters can be supplied from a file or on the command line. The file contents is in JSON format and
	must be in the same directory as the binary using the core services.
		Example: Business Service X (Path: ../BusinessService/X) is using core logging, so the init file must be in the
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
	"github.com/sy-consulting/golang/core/cSystemInfo"
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
func New(application, environment string, si cSystemInfo.SystemInfoIF) (coreManagerPtr *Manager) {

	coreManagerPtr = &Manager{
		application: application,
		environment: environment,
	}

	// Start Work here
	x := si.GetIP
	fmt.Println(x)

	return
}
