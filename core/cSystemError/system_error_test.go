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
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	if x == nil {
		t.Error("SystemError.New failed to return an object")
	}
}

func TestErrItemNotPopulated_100100(t *testing.T) {
	x := New(TEST_APPLICATION, TEST_ENVIRONMENT, TEST_INTERNALIP)
	myError := x.ErrItemNotPopulated_100100(TEST_APPLICATION)
	if myError == nil {
		t.Error("SystemError.ErrItemNotPopulated_100100 failed to return an object")
	}
	if myError.ErrorCode != NOT_POPULATED {
		t.Error("SystemError.ErrItemNotPopulated_100100 has the wrong error code")
	}
	if myError.LineNumber == 0 {
		t.Error("SystemError.ErrItemNotPopulated_100100 is missing the line number")
	}
	if myError.FileName == "" {
		t.Error("SystemError.ErrItemNotPopulated_100100 is missing the FileName")
	}
	if myError.ErrorMsg == "" {
		t.Error("SystemError.ErrItemNotPopulated_100100 is missing the error message")
	}
}
