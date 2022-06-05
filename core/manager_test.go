/*
NOTES:
    This only tests manager, not any of the packages contained in core.
*/
package core

import (
	"fmt"
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION = "application"
	TEST_ENVIRONMENT = "environment"
)

func TestCoreNew(t *testing.T) {
	type arguments struct {
		application     string
		environment     string
		internalIPFound bool
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
				application:     TEST_APPLICATION,
				environment:     TEST_ENVIRONMENT,
				internalIPFound: true,
			},
			desiredResult: true,
		},
		{
			name: "negative-case: application and environment are missing",
			arguments: arguments{
				application:     "",
				environment:     "",
				internalIPFound: false,
			},
			desiredResult: false,
		},
		{
			name: "negative-case: application is missing and environment is populated",
			arguments: arguments{
				application:     "",
				environment:     TEST_ENVIRONMENT,
				internalIPFound: false,
			},
			desiredResult: false,
		},
		{
			name: "negative-case: application is populated and environment is missing",
			arguments: arguments{
				application:     TEST_APPLICATION,
				environment:     "",
				internalIPFound: false,
			},
			desiredResult: false,
		},
	}
	for _, ts := range tests {
		actualResult = true
		t.Run(ts.name, func(t *testing.T) {
			var (
				x = New(ts.arguments.application, ts.arguments.environment)
			)
			if x.application != TEST_APPLICATION {
				actualResult = false
				errorMessages = append(errorMessages, fmt.Sprintf("New(application, environment string) Parameters passed '%v' and '%v' / application is not valid", ts.arguments.application, ts.arguments.environment))
			}
			if x.environment != TEST_ENVIRONMENT {
				actualResult = false
				errorMessages = append(errorMessages, fmt.Sprintf("New(application, environment string) Parameters passed '%v' and '%v' / environment is not valid", ts.arguments.application, ts.arguments.environment))
			}
			if x.internalIP == nil {
				actualResult = false
				errorMessages = append(errorMessages, fmt.Sprintf("New(application, environment string) Parameters passed '%v' and '%v' / internalIP (%v) on system is not valid", ts.arguments.application, ts.arguments.environment, x.internalIP))
			}
		})
		if actualResult != ts.desiredResult {
			for _, message := range errorMessages {
				t.Error(message)
			}
		}
	}

}
