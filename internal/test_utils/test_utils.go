package test_utils

import (
	"testing"
)

func EnsureNotError(t *testing.T, error error) {
	if error != nil {
		t.Errorf("error != nil, error is \"%s\"", error)
	}
}

func EnsureEqual(t *testing.T, expected any, actual any) {
	if expected != actual {
		t.Errorf("'%s' != '%s'", expected, actual)
	}
}
