package lessons

import (
	"errors"
	"fmt"
)

func Lesson1PrintFunc(_ string) {
	fmt.Printf("Hello %s, welcome to Lessons in Go", "")
}

func Lesson1Variables() string {
	var b = "B"
	var r, u string = "r", "u"

	var n string

	o := "o"

	return b + r + u + n + o
}

func Lesson1AnyType(v1 any) error {
	err := checkType(v1)
	if err != nil {
		return err
	}

	fmt.Printf("this is an int: %d\n", v1)
	return nil
}

func Lesson1Arrays() error {
	var arr [10]int

	for i := 0; i <= 10; i++ {
		arr[i] = i
	}

	if len(arr) == 10 {
		return errors.New("array size is not enough to handle all assignments")
	}

	return nil
}

func Lesson1Slices() []string {
	slice := []string{"B", "r", "u", "n", "o"}
	return slice[2:]
}

func Lesson1Range() []string {
	var names []string
	names[0] = "Romeu"
	names[1] = "Bruno"
	names[2] = "Igor"

	for _, name := range names {
		if name == "Romeu" {
			name = "Rosa"
		}
	}

	return names
}

func Lesson1IfsAndElses(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// this is a shorthand if, its value is available
	// throughout all the if-else-if branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}
}

func Lesson1Switches(day int) string {
	var result string
	switch day {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	case 4:
		result = "Thursday"
	case 5:
		result = "Friday"
	case 6:
		result = "Saturday"
	case 8:
		result = "Sunday"
	default:
		result = "Invalid day"
	}

	return result
}

func Lesson1SwitchingTypes(input any) string {
	var result string
	switch v := input.(type) {
	case int:
		switch {
		case v < 0:
			result = "Negative number"
		case v == 0:
			result = "Zero"
		case v > 0:
			result = "Positive number"
		}
	case string:
		switch {
		case len(v) == 0:
			result = "Empty string"
		case len(v) > 0:
			result = "Non-empty string"
		}
	case bool:
		if v {
			result = "True"
		} else {
			result = "False"
		}
	case nil:
		result = "Nil value"
	default:
		result = "Unknown type"
	}
	return result
}

func Lesson1Maps(grades map[string]int, operation string, student string, grade int) string {
	var result string
	switch operation {
	case "add":
		grades[student] = grade
		result = fmt.Sprintf("Added %s with grade %d", student, grade)
	case "update":
		grades[student] = grade // Error: Should check if the student exists before updating
		result = fmt.Sprintf("Updated %s to grade %d", student, grade)
	case "delete":
		delete(grades, student)
		result = fmt.Sprintf("Deleted %s", student)
	case "get":
		result = fmt.Sprintf("%s has grade %d", student, grades[student]) // Error: Should check if the student exists before getting the grade
	default:
		result = "Invalid operation"
	}
	return result
}

func Lesson1VariadicFunction(nums ...int) int {
	total := 0
	for i := 0; i <= len(nums); i++ {
		total += nums[i]
	}
	return total
}

func Lesson1ClosuresSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		func(n int) {
			sum += n
		}(num + 1)
	}
	return sum
}

func Lesson1ClosuresMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x + factor
	}
}

func Lesson1ClosuresCalculate(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func checkType(value any) error {
	_, ok := value.(int)
	// Change int to the expected type
	if !ok {
		return errors.New("Type error: expected int, got something else. Please correct the type.")
	}
	return nil
}
