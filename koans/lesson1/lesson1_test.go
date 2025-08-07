package lesson1

import (
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// TestHello verifies that the Hello function returns the exact expected output.
func TestHello(t *testing.T) {
	expected := "Hello, World!"
	testutils.Compare(t, expected, Hello(), "The Hello function should return 'Hello, World!'")
}
