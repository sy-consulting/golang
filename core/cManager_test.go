package core

import (
	"fmt"
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION = "application"
	TEST_ENVIRONMENT = "environment"
	MOCK             = true
	MOCK_PASSED      = "core.SystemInfoMock{}"
)

func TestNew(t *testing.T) {

	type arguments struct {
		application string
		environment string
	}

	var (
		errorMessages []string
	)

	tests := []struct {
		name          string
		arguments     arguments
		desiredResult bool
	}{
		{
			name: "positive-case: application and environment are populated",
			arguments: arguments{
				application: TEST_APPLICATION,
				environment: TEST_ENVIRONMENT,
			},
			desiredResult: true,
		},
		{
			name: "negative-case: application and environment are missing",
			arguments: arguments{
				application: "",
				environment: "",
			},
			desiredResult: false,
		},
		{
			name: "negative-case: application is missing and environment is populated",
			arguments: arguments{
				application: "",
				environment: TEST_ENVIRONMENT,
			},
			desiredResult: false,
		},
		{
			name: "negative-case: application is populated and environment is missing",
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
			if err != nil {
				errorMessages = append(errorMessages, fmt.Sprintf("Parameters passed to New('%v', '%v', '%v') are not valid", ts.arguments.application, ts.arguments.environment,
					"core.SystemInfoMock{}"))
			}
		})
	}
}
