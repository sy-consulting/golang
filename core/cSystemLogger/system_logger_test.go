/*
	This test is a manual check of the results. Capturing the StdOut didn't work so the results could be checked programmatically.
*/

package cSystemLogger

import (
	"testing"
)

//goland:noinspection ALL
const (
	TEST_APPLICATION = "application"
	TEST_ENVIRONMENT = "environment"
	TEST_INTERNALIP  = "192.0.0.0"
)

func TestNew(t *testing.T) {
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	if x == nil {
		t.Error("SystemLogger.New failed to return an object")
	}
}

func TestDebugOn(t *testing.T) {
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	x.TurnDebugOn()
	if !x.IsDebugOn() {
		t.Error("SystemLogger.TurnDebugOn is not working")
	}
}

func TestDebugOff(t *testing.T) {
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	x.TurnDebugOff()
	if x.IsDebugOn() {
		t.Error("SystemLogger.TurnDebugOff is not working")
	}
}

func TestLogPanic(t *testing.T) {
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	x.LogPanic(nil)
}

func TestLogFatal(t *testing.T) {
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	x.LogFatal(nil)
}
