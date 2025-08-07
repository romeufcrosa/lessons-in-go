package lesson2

import (
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

func TestLesson2BasicSyntax(t *testing.T) {
	expected := "Hello, the answer is 42, pi is 3.14, truth is true!"
	actual := Lesson2BasicSyntax()
	hint := "Ensure that 'pi' is assigned 3.14 (not 3.15) and 'truth' is set to true (not false)."
	testutils.Compare(t, expected, actual, hint)
}

// TestArithmeticOps verifies that ArithmeticOps returns the correct result.
// Expected output: "Sum is 12, Product is 20, Division is 5.00, Divisibility is true"
func TestArithmeticOps(t *testing.T) {
	expected := "Sum is 12, Product is 20, Division is 5.00, Divisibility is true"
	actual := Lesson2ArithmeticOps()
	hint := "Ensure arithmetic operations are correct and that the divisibility check uses '==' instead of '!='."
	testutils.Compare(t, expected, actual, hint)
}

func TestLesson2Variables(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			testutils.Compare(t, nil, r, "arrays have limits")
		}
	}()
	output := Lesson2Variables()
	testutils.Compare(t, "Bruno", output, "many ways of declaring a variable, make sure they have the right values")
}

func TestLesson2AnyType(t *testing.T) {
	output := Lesson2AnyType("1")
	testutils.Compare(t, nil, output, "any is useful, but make sure you assert the right type")
}
