package api

import (
	"testing"

	"github.com/lsh-0/ppp-go/internal/test_utils"
)

func TestParseContentType(t *testing.T) {
	actual_mime, actual_version, error := ParseContentType("application/problem+json; version=1")
	expected_mime := "application/problem+json"
	expected_version := 1
	test_utils.EnsureNotError(t, error)
	test_utils.EnsureEqual(t, expected_mime, actual_mime)
	test_utils.EnsureEqual(t, expected_version, actual_version)
}

func TestParseContentTypeNoVersionParameter(t *testing.T) {
	actual_mime, _, error := ParseContentType("application/problem+json")
	expected_mime := "application/problem+json"
	test_utils.EnsureNotError(t, error)
	test_utils.EnsureEqual(t, expected_mime, actual_mime)
}
