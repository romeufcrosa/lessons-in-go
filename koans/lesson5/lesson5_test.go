package lesson5

import (
	"testing"

	"github.com/romeufcrosa/lessons-in-go/koans/lesson5/calc"
	"github.com/romeufcrosa/lessons-in-go/testutils"
)

func TestLesson5VariadicFunctions(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 6, 7}, 22},
		{[]int{}, 0},
	}

	for _, test := range tests {
		defer func() {
			if r := recover(); r != nil {
				testutils.PrintFailure(t, test.expected, r.(error).Error(), "variadic or not, a slice is still a slice")
			}
		}()
		result := Lesson5VariadicFunction(test.input...)
		testutils.Compare(t, test.expected, result, "check the error")
	}
}

func TestLesson5Closures(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expectedSum := 15
	actualSum := Lesson5ClosuresSum(numbers)

	testutils.Compare(t, expectedSum, actualSum, "")

	double := Lesson5ClosuresMultiplier(2)
	triple := Lesson5ClosuresMultiplier(3)

	expectedDouble := 10
	actualDouble := double(5)
	testutils.Compare(t, expectedDouble, actualDouble, "")

	expectedTriple := 15
	actualTriple := triple(5)
	testutils.Compare(t, expectedTriple, actualTriple, "")

	_, err := Lesson5ClosuresCalculate(10, 0)
	testutils.Compare(t, nil, err.Error(), "")

	expectedResult := 100
	actualResult, _ := Lesson5ClosuresCalculate(20, 5)
	testutils.Compare(t, expectedResult, actualResult, "")
}

// TestAddSubtract checks that AddSubtract returns the correct summation and difference.
// For inputs 5 and 3, the expected results are: sum = 8 and difference = 2.
func TestLesson5AddSubtract(t *testing.T) {
	sum, diff := calc.AddSubtract(5, 3)
	expected := []int{8, 2}
	actual := []int{sum, diff}
	hint := "Ensure that AddSubtract returns the correct sum and difference. For inputs (5, 3), the difference should be computed as 5-3, not 5-3+1."
	testutils.Compare(t, expected, actual, hint)
}

// TestSumAll checks that SumAll correctly returns the sum of its arguments.
// For inputs 5, 3, and 10, the expected sum is 18.
func TestLesson5SumAll(t *testing.T) {
	result := calc.SumAll(5, 3, 10)
	expected := 18
	hint := "Initialize the sum to 0 in SumAll instead of 1 so that the function correctly calculates the total."
	testutils.Compare(t, expected, result, hint)
}

// TestClosureDemo verifies that the closure correctly doubles the input number.
// For an input of 5, the expected result is 10.
func TestLesson5ClosureDemo(t *testing.T) {
	result := Lesson5ClosureDemo(5)
	expected := 10
	hint := "The closure in ClosureDemo should double the number (2*x). Change the multiplier from 3 to 2."
	testutils.Compare(t, expected, result, hint)
}
