package cSystemLogger

import (
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION = "application"
	TEST_ENVIRONMENT = "environment"
	TEST_INTERNALIP  = "192.0.0.0"
	//TEST_SPLIT_ERRORS      = true
	//TEST_DONT_SPLIT_ERRORS = false
)

func TestNew(t *testing.T) {

	type arguments struct {
		application string
		environment string
		internalIP  string
		//splitErrors  bool
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
				//splitErrors:  TEST_DONT_SPLIT_ERRORS,
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
			x.TurnDebugOn()
			x.TurnDebugOff()
			x.TurnDebugOn()
		})
		if actualResult != ts.desiredResult {
			for _, message := range errorMessages {
				t.Error(message)
			}
		}
	}
}
