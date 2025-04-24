package lessons

import (
	"testing"
)

func TestLesson1PrintFunc(t *testing.T) {
	output := captureOutput(func() { Lesson1PrintFunc("Bruno") })
	compare(t, "Hello Bruno, welcome to Lessons in Go", output, "underscore hides things")
}

func TestLesson1Variables(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			compare(t, nil, r, "arrays have limits")
		}
	}()
	output := Lesson1Variables()
	compare(t, "Bruno", output, "many ways of declaring a variable, make sure they have the right values")
}

func TestLesson1AnyType(t *testing.T) {
	output := Lesson1AnyType("1")
	compare(t, nil, output, "any is useful, but make sure you assert the right type")
}

func TestLesson1Array(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			compare(t, nil, r, "arrays have limits")
		}
	}()
	err := Lesson1Arrays()
	compare(t, nil, err, "you have found an alternative path, but try modifying the bounds first")
}

func TestLesson1Slices(t *testing.T) {
	output := Lesson1Slices()
	compare(t, "uno", output, "a slice of a slice is still a slice")
}

func TestLesson1Range(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			compare(t, []string{"Rosa", "Bruno", "Igor"}, r, "declaration is not initialization")
		}
	}()
	output := Lesson1Range()
	compare(t, []string{"Rosa", "Bruno", "Igor"}, output, "range is useful, but not in this scenario. do you know why?")
}

func TestLesson1IfsAndElses(t *testing.T) {
	output := captureOutput(func() { Lesson1IfsAndElses(10) })
	compare(t, "10 is even\n10 has two digits\n", output, "if branches can get confusing, maybe there is also one missing?")
}

func TestLesson1Switch(t *testing.T) {
	output := Lesson1Switches(1)
	compare(t, "Monday", output, "Monday is the worst")
	output = Lesson1Switches(8)
	compare(t, "Invalid day", output, "that is weird, weeks don't have 8 days")
}

func TestLesson1SwitchingTypes(t *testing.T) {
	compare(t, "Negative number", Lesson1SwitchingTypes(-5), "")
	compare(t, "Zero", Lesson1SwitchingTypes(0), "")
	compare(t, "Positive number", Lesson1SwitchingTypes(10), "")
	compare(t, "Empty string", Lesson1SwitchingTypes(""), "")
	compare(t, "Non-empty string", Lesson1SwitchingTypes("hello"), "")
	compare(t, "True", Lesson1SwitchingTypes(true), "")
	compare(t, "False", Lesson1SwitchingTypes(false), "")
	compare(t, "Nil value", Lesson1SwitchingTypes(nil), "")
	compare(t, "Floating-point number", Lesson1SwitchingTypes(3.14), "in the old days computers had a dedicated processor just to make FP calculations")
	compare(t, "Complex number", Lesson1SwitchingTypes(1+2i), "oops, I think I forgot complex numbers")
}

func TestLesson1Maps(t *testing.T) {
	grades := map[string]int{
		"Bruno": 15,
		"Igor":  12,
	}

	// Test adding a new student
	result := Lesson1Maps(grades, "add", "Romeu", 10)
	compare(t, "Added Romeu with grade 10", result, "")

	// Test updating an existing student
	result = Lesson1Maps(grades, "update", "Bruno", 16)
	compare(t, "Updated Bruno to grade 16", result, "")

	// Test updating a non-existing student (should fail)
	result = Lesson1Maps(grades, "update", "Margarida", 18)
	compare(t, "Student Margarida not found", result, "I don't remember creating Margarida")

	// Test getting a non-existing student's grade (should fail)
	result = Lesson1Maps(grades, "get", "Margarida", 0)
	compare(t, "Student Margarida not found", result, "still don't think Margarida actually exists")

}

func TestLesson1VariadicFunctions(t *testing.T) {
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
				printFailure(t, test.expected, r.(error).Error(), "variadic or not, a slice is still a slice")
			}
		}()
		result := Lesson1VariadicFunction(test.input...)
		compare(t, test.expected, result, "check the error")
	}
}

func TestLesson1Closures(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expectedSum := 15
	actualSum := Lesson1ClosuresSum(numbers)

	compare(t, expectedSum, actualSum, "")

	double := Lesson1ClosuresMultiplier(2)
	triple := Lesson1ClosuresMultiplier(3)

	expectedDouble := 10
	actualDouble := double(5)
	compare(t, expectedDouble, actualDouble, "")

	expectedTriple := 15
	actualTriple := triple(5)
	compare(t, expectedTriple, actualTriple, "")

	_, err := Lesson1ClosuresCalculate(10, 0)
	compare(t, nil, err.Error(), "")

	expectedResult := 100
	actualResult, _ := Lesson1ClosuresCalculate(20, 5)
	compare(t, expectedResult, actualResult, "")
}
