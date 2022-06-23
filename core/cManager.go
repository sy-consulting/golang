/*
 */
package core

import (
	"github.com/sy-consulting/golang/core/cSystemError"
	"github.com/sy-consulting/golang/core/cSystemInfo"
	"github.com/sy-consulting/golang/core/cSystemLogger"
)

//goland:noinspection ALL
const (
	CORE_APPLICATION = "APPLICATION"
	CORE_ENVIRONMENT = "ENVIRONMENT"
)

type Manager struct {
	application  string
	environment  string
	internalIP   string
	systemInfo   cSystemInfo.SystemInfoIF
	systemLogger *cSystemLogger.SystemLogger
	systemError  *cSystemError.SystemError
}

// Core.New
func New(application, environment string, mock bool) (cmPtr *Manager, myError *cSystemError.Error) {

	cmPtr = &Manager{
		application: application,
		environment: environment,
		internalIP:  "",
	}

	cmPtr.systemError = cSystemError.New(application, environment, cmPtr.internalIP)
	if application == "" {
		return nil, cmPtr.systemError.ErrItemNotPopulated_100100(CORE_APPLICATION)
	}
	if environment == "" {
		return nil, cmPtr.systemError.ErrItemNotPopulated_100100(CORE_ENVIRONMENT)
	}

	cmPtr.systemInfo = cSystemInfo.SystemInfoIF(cSystemInfo.SystemInfo{})
	if mock {
		cmPtr.systemInfo = cSystemInfo.SystemInfoIF(cSystemInfo.SystemInfoMock{})
	}

	cmPtr.internalIP = cmPtr.systemInfo.GetIP()
	cmPtr.systemLogger = cSystemLogger.New(application, environment, cmPtr.internalIP)

	return
}
