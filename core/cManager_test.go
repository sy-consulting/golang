package core

import (
	"fmt"
	core "github.com/sy-consulting/golang/core/cSystemInfo"
	"testing"
)

func TestNew(t *testing.T) {
	cmPtr := New("application_test", "environment_test", core.SystemInfoIF(core.SystemInfoMock{}))
	fmt.Printf("Mock IP: %v\n", cmPtr.internalIP)
	cmPtr = New("application_test", "environment_test", core.SystemInfoIF(core.SystemInfo{}))
	fmt.Printf("Real IP: %v\n", cmPtr.internalIP)
}
