/*
core oversees all the shared services, such as logging, errors, messaging, etc.
	The input can be from a file or on the command line. The init file is in JSON format and must be in the same
	directory as the binary using the core services.
		Example: Business Service X (Path: ../BusinessService/X) is using core logging, so the init file must be in the
		../BusinessService directory.

	All errors are in JSON format.
	JSON Fields:
		application      Blah Blah

Notes:
	None
*/
package core

const (
// Constants go here
)

type Manager struct {
	application string
	environment string
}

// Starts logging
func New(application, environment, logLevel string) (coreManagerPtr *Manager) {
	// Initialize the values for Nats Manager
	coreManagerPtr = &Manager{
		application: application,
		environment: environment,
	}
	return coreManagerPtr
}
