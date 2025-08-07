package lesson6

import (
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// TestGreet ensures that the Greet method returns the correct string.
func TestLesson6Greet(t *testing.T) {
	alice := Person{Name: "Alice", Age: 30}
	expected := "Hello, my name is Alice and I am 30 years old."
	actual := alice.Greet()
	hint := "Ensure that the Greet method correctly formats the greeting using the Person's Name and Age."
	testutils.Compare(t, expected, actual, hint)
}

// TestBirthday verifies that the Birthday method increments the Age.
// Due to the bug, the Age remains unchanged—fix it by using a pointer receiver.
func TestLesson6Birthday(t *testing.T) {
	bob := Person{Name: "Bob", Age: 40}
	bob.Birthday()
	expected := 41
	actual := bob.Age
	hint := "Birthday should increment the Person's Age. If the Age is not updated, change the method receiver to a pointer."
	testutils.Compare(t, expected, actual, hint)
}

// TestCorrectBirthday shows that using a pointer receiver correctly updates Age.
// This serves as an example for correct behavior.
func TestLesson6CorrectBirthday(t *testing.T) {
	charlie := Person{Name: "Charlie", Age: 50}
	charlie.CorrectBirthday()
	expected := 51
	actual := charlie.Age
	hint := "CorrectBirthday should properly increment Age using a pointer receiver."
	testutils.Compare(t, expected, actual, hint)
}

func TestLesson6Pointers(t *testing.T) {
	// Test pointer swap
	a, b := 1, 2
	Lesson6PointersSwap(&a, &b)
	testutils.Compare(t, 2, a, "Expected a=2 after swap")
	testutils.Compare(t, 1, b, "Expected b=1 after swap")

	// Test SetTo42
	x := 10
	Lesson6PointersSetTo42(&x)
	testutils.Compare(t, 42, x, "Expected x=42 after SetTo42")

	// Test nil pointer check
	var p *int
	testutils.Compare(t, true, Lesson6PointersNilPointerCheck(p), "Expected true for nil pointer")

	x = 5
	testutils.Compare(t, false, Lesson6PointersNilPointerCheck(&x), "Expected false for non-nil pointer")
}
