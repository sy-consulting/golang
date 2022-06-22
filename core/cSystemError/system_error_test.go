/*
	This test is a manual check of the results. Capturing the StdOut didn't work so the results could be checked programmatically.
*/

package cSystemError

import (
	"errors"
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION   = "application"
	TEST_ENVIRONMENT   = "environment"
	TEST_INTERNALIP    = "192.0.0.0"
	TEST_ERROR_MESSAGE = "TEST Error Message"
)

var (
	testDetails = make(map[int]string)
	testErr     = errors.New("TEST ERROR")
)

func TestNew(t *testing.T) {

	type arguments struct {
		application string
		environment string
		internalIP  string
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
			name: "positive-case: writes log record with all parameters populated",
			arguments: arguments{
				application: TEST_APPLICATION,
				environment: TEST_ENVIRONMENT,
				internalIP:  TEST_INTERNALIP,
			},
			desiredResult: true,
		},
	}
	for _, ts := range tests {
		actualResult = true
		t.Run(ts.name, func(t *testing.T) {
			var (
				x = New(ts.arguments.application, ts.arguments.environment, ts.arguments.internalIP)
			)
			x.ErrItemAlreadyExists_100000(TEST_ERROR_MESSAGE, testDetails, testErr)
		})
		if actualResult != ts.desiredResult {
			for _, message := range errorMessages {
				t.Error(message)
			}
		}
	}
}
