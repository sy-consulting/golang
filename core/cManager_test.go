package core

import (
	"fmt"
	core "github.com/sy-consulting/golang/core/cSystemInfo"
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION = "application"
	TEST_ENVIRONMENT = "environment"
	MOCK_PASSED      = "core.SystemInfoMock{}"
)

func TestNew(t *testing.T) {

	type arguments struct {
		application string
		environment string
	}

	var (
		errorMessages []string
		actualResult  bool
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
		actualResult = true
		t.Run(ts.name, func(t *testing.T) {
			var (
				x = New(ts.arguments.application, ts.arguments.environment, core.SystemInfoIF(core.SystemInfoMock{}))
			)
			if x.application != TEST_APPLICATION {
				actualResult = false
				errorMessages = append(errorMessages, fmt.Sprintf("Parameters passed to New('%v', '%v', '%v') and the application is not valid", ts.arguments.application, ts.arguments.environment,
					"core.SystemInfoMock{}"))
			}
			if x.environment != TEST_ENVIRONMENT {
				actualResult = false
				errorMessages = append(errorMessages, fmt.Sprintf("Parameters passed to New('%v', '%v', '%v') and the environment is not valid", ts.arguments.application, ts.arguments.environment,
					"core.SystemInfoMock{}"))
			}
		})
		if actualResult != ts.desiredResult {
			for _, message := range errorMessages {
				t.Error(message)
			}
		}
	}
}
