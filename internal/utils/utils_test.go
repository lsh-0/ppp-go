package utils

import (
	"testing"

	"github.com/lsh-0/ppp-go/internal/test_utils"
)

func TestToJSON(t *testing.T) {
	actual := ToJSON("foo")
	expected := "\"foo\""
	test_utils.EnsureEqual(t, expected, actual)
}

func TestFromJSON(t *testing.T) {
	actual := FromJSON[string]([]byte("\"foo\""))
	expected := "foo"
	test_utils.EnsureEqual(t, expected, actual)
}
