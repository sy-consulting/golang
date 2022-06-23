package core

import (
	"github.com/sy-consulting/golang/core/cSystemError"
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION = "application"
	TEST_ENVIRONMENT = "environment"
	MOCK             = true
	MOCK_PASSED      = "core.SystemInfoMock{}"
)

func TestNewApplication(t *testing.T) {

	type arguments struct {
		application string
		environment string
	}

	tests := []struct {
		name          string
		arguments     arguments
		desiredResult bool
	}{
		{
			name: "positive-case: application parameter is populated",
			arguments: arguments{
				application: TEST_APPLICATION,
				environment: TEST_ENVIRONMENT,
			},
			desiredResult: true,
		},
		{
			name: "negative-case: application parameter are missing",
			arguments: arguments{
				application: "",
				environment: TEST_ENVIRONMENT,
			},
			desiredResult: false,
		},
		{
			name: "positive-case: environment parameter is populated",
			arguments: arguments{
				application: TEST_APPLICATION,
				environment: TEST_ENVIRONMENT,
			},
			desiredResult: true,
		},
		{
			name: "negative-case: environment parameter are missing",
			arguments: arguments{
				application: TEST_APPLICATION,
				environment: "",
			},
			desiredResult: false,
		},
	}
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			_, err := New(ts.arguments.application, ts.arguments.environment, MOCK)
			if err != nil && err.ErrorCode != cSystemError.NOT_POPULATED {
				t.Error(ts.name + " check failed")
			}
		})
	}
}
