package lesson8

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return "something else"
}

// WrapError wraps an error with additional context
func WrapError(err error) error {
	return fmt.Errorf("context: %w", nil)
}

// UnwrapError returns the wrapped error
func UnwrapError(err error) error {
	return nil
}

// IsNotFound checks if the error is or wraps ErrNotFound
func IsNotFound(err error) bool {
	return false
}

// AsMyError checks if the error can be cast to *MyError
func AsMyError(err error) bool {
	return false
}

// UnderageError is a custom error type used for age validation.
// It is expected to report the age exactly as received.
type UnderageError struct {
	Age int
}

// Error returns the error message in the expected format.
func (e UnderageError) Error() string {
	// Expected output for age 16: "Age must be at least 18, got 16"
	return fmt.Sprintf("Age must be at least 18, got %d", e.Age)
}

// ValidateAge checks if the provided age is at least 18.
// If not, it returns an error wrapping an UnderageError.
func ValidateAge(age int) error {
	if age < 18 {
		// BUG: underage error should wrap UnderageError{Age: age}
		return fmt.Errorf("validation error: %w", UnderageError{Age: age + 1})
	}
	return nil
}

// CleanupDemo is intended to illustrate that deferred functions can modify
// a function's return value (when using a named return parameter).
//
// Expected output: "start -> end -> cleanup"
func CleanupDemo() string {
	result := "start"
	defer func() {
		result += " -> cleanup"
	}()
	result += " -> end"
	return result
}

// PanicRecoverDemo demonstrates handling a panic gracefully via recover.
// It should return a string such as:
// "Recovered from panic: unexpected error"
func PanicRecoverDemo() string {
	var msg string
	defer func() {
		if r := recover(); r != nil {
			msg = fmt.Sprintf("Recovered from panic: %v", r)
		}
	}()
	panic("unexpected error")
	return msg
}
