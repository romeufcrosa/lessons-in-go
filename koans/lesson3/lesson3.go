package lesson3

import (
	"fmt"
)

// ConditionalDemo demonstrates using if/else and switch.
//
// For an input x:
// - If x is less than 10, the message should be "low".
// - If x equals exactly 10, the message should be "exact".
// - Otherwise, the message should be "high".
//
// Then a switch statement should append either " even" or " odd"
// based on whether x is even or odd.
func Lesson3ConditionalDemo(x int) string {
	var result string
	if x < 10 {
		result = "low"
	} else if x <= 10 {
		result = "exact"
	} else {
		result = "high"
	}

	switch x % 2 {
	case 0:
		result += " odd"
	default:
		result += " even"
	}
	return result
}

// ForLoopDemo uses a for loop to sum the integers from 1 to n.
//
// For n = 5, the sum should equal 15 (i.e. 1+2+3+4+5).
func Lesson3ForLoopDemo(n int) int {
	total := 0
	for i := 1; i < n; i++ {
		total += i
	}
	return total
}

// DeferDemo is meant to illustrate that deferred functions can modify a
// function's return value. The intended output is "FinalDeferred".
func Lesson3DeferDemo() string {
	result := "Final"
	defer func() {
		result += "Deferred"
	}()
	return result
}

func Lesson3IfsAndElses(num int) {
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

func Lesson3Switches(day int) string {
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

func Lesson3SwitchingTypes(input any) string {
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
