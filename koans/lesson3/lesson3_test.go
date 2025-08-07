package lesson3

import (
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// TestConditionalDemo checks that for x = 10, ConditionalDemo returns "exact even".
func TestConditionalDemo(t *testing.T) {
	expected := "exact even"
	actual := Lesson3ConditionalDemo(10)
	hint := "For input 10, ensure the if/else block uses 'x == 10' to yield 'exact' and the switch appends ' even' for even numbers."
	testutils.Compare(t, expected, actual, hint)
}

// TestForLoopDemo verifies that ForLoopDemo sums numbers from 1 to n correctly.
// For n = 5, the expected sum is 15.
func TestForLoopDemo(t *testing.T) {
	expected := 15
	actual := Lesson3ForLoopDemo(5)
	hint := "Ensure the for loop includes the upper bound (use 'i <= n') so that all numbers from 1 to n are summed."
	testutils.Compare(t, expected, actual, hint)
}

// TestDeferDemo checks that DeferDemo returns "FinalDeferred".
// Hint: use a named return parameter so that the deferred function can modify the returned value.
func TestDeferDemo(t *testing.T) {
	expected := "FinalDeferred"
	actual := Lesson3DeferDemo()
	hint := "To allow the deferred function to update the return value, change the function signature to use a named return parameter."
	testutils.Compare(t, expected, actual, hint)
}

func TestLesson3IfsAndElses(t *testing.T) {
	output := testutils.CaptureOutput(func() { Lesson3IfsAndElses(10) })
	testutils.Compare(t, "10 is even\n10 has two digits\n", output, "if branches can get confusing, maybe there is also one missing?")
}

func TestLesson3Switch(t *testing.T) {
	output := Lesson3Switches(1)
	testutils.Compare(t, "Monday", output, "Monday is the worst")
	output = Lesson3Switches(8)
	testutils.Compare(t, "Invalid day", output, "that is weird, weeks don't have 8 days")
}

func TestLesson3SwitchingTypes(t *testing.T) {
	testutils.Compare(t, "Negative number", Lesson3SwitchingTypes(-5), "")
	testutils.Compare(t, "Zero", Lesson3SwitchingTypes(0), "")
	testutils.Compare(t, "Positive number", Lesson3SwitchingTypes(10), "")
	testutils.Compare(t, "Empty string", Lesson3SwitchingTypes(""), "")
	testutils.Compare(t, "Non-empty string", Lesson3SwitchingTypes("hello"), "")
	testutils.Compare(t, "True", Lesson3SwitchingTypes(true), "")
	testutils.Compare(t, "False", Lesson3SwitchingTypes(false), "")
	testutils.Compare(t, "Nil value", Lesson3SwitchingTypes(nil), "")
	testutils.Compare(t, "Floating-point number", Lesson3SwitchingTypes(3.14), "in the old days computers had a dedicated processor just to make FP calculations")
	testutils.Compare(t, "Complex number", Lesson3SwitchingTypes(1+2i), "oops, I think I forgot complex numbers")
}
