package lesson8

import (
	"errors"
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

func TestLesson2Errors(t *testing.T) {
	// Test WrapError and UnwrapError
	wrapped := WrapError(ErrNotFound)
	testutils.Compare(t, true, errors.Is(wrapped, ErrNotFound), "WrapError should wrap ErrNotFound")
	testutils.Compare(t, ErrNotFound, UnwrapError(wrapped), "UnwrapError should return the original error")

	// Test IsNotFound
	testutils.Compare(t, true, IsNotFound(wrapped), "IsNotFound should return true for wrapped ErrNotFound")

	// Test AsMyError
	myErr := &MyError{Msg: "custom"}
	wrappedMyErr := WrapError(myErr)
	testutils.Compare(t, true, AsMyError(wrappedMyErr), "AsMyError should return true for wrapped MyError")
}

// TestValidateAge verifies the custom error from ValidateAge.
// When age is 16, the expected error message is:
// "validation error: Age must be at least 18, got 16"
// The intentional bug causes it to report "got 17".
func TestValidateAge(t *testing.T) {
	err := ValidateAge(16)
	expectedErrMsg := "validation error: Age must be at least 18, got 16"
	if err == nil {
		t.Errorf("Expected an error for age 16, but got nil")
	} else {
		hint := "Ensure that when returning the error, you wrap the UnderageError with the correct age (do not add 1 to the age)."
		testutils.Compare(t, expectedErrMsg, err.Error(), hint)
	}

	err = ValidateAge(20)
	_ = "For age >= 18, ValidateAge should return nil without error."
	if err != nil {
		t.Errorf("Expected no error for age 20, but got: %v", err)
	} else {
		t.Log("ValidateAge passes for age 20.")
	}
}

// TestCleanupDemo verifies that CleanupDemo returns the expected string.
// Expected output: "start -> end -> cleanup"
func TestCleanupDemo(t *testing.T) {
	expected := "start -> end -> cleanup"
	actual := CleanupDemo()
	hint := "To allow the deferred function to modify the return value, use a named return parameter instead of a local variable."
	testutils.Compare(t, expected, actual, hint)
}

// TestPanicRecoverDemo checks that PanicRecoverDemo gracefully recovers from panic.
// Expected output: "Recovered from panic: unexpected error"
func TestPanicRecoverDemo(t *testing.T) {
	expected := "Recovered from panic: unexpected error"
	actual := PanicRecoverDemo()
	hint := "PanicRecoverDemo should use a named return parameter so that the deferred recover block can set the returned value."
	testutils.Compare(t, expected, actual, hint)
}
