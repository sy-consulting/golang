/*
 */
package core

import (
	"github.com/sy-consulting/golang/core/cSystemInfo"
	"github.com/sy-consulting/golang/core/cSystemLogger"
)

const (
	CORE_MESSAGES = "MESSAGES"
	CORE_DATABASE = "DATABASE"
)

type Manager struct {
	application  string
	environment  string
	internalIP   string
	systemInfo   cSystemInfo.SystemInfoIF
	systemLogger *cSystemLogger.SystemLogger
}

// Core.New
func New(application, environment string, mock bool) (cmPtr *Manager) {

	cmPtr = &Manager{
		application: application,
		environment: environment,
		internalIP:  "",
	}

	cmPtr.systemInfo = cSystemInfo.SystemInfoIF(cSystemInfo.SystemInfo{})

	if mock {
		cmPtr.systemInfo = cSystemInfo.SystemInfoIF(cSystemInfo.SystemInfoMock{})
	}

	cmPtr.internalIP = cmPtr.systemInfo.GetIP()
	cmPtr.systemLogger = cSystemLogger.New(application, environment, cmPtr.internalIP)
	cmPtr.systemLogger.LogMessage("This is a test")

	return
}
