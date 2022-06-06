package core

import (
	core "github.com/sy-consulting/golang/core/cSystemInfo"
	"testing"
)

func TestNew(t *testing.T) {
	New("application_test", "environment_test", core.SystemInfoIF(core.SystemInfoMock{}))
	New("application_test", "environment_test", core.SystemInfoIF(core.SystemInfo{}))
}
