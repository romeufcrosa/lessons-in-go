package lesson2

import (
	"errors"
	"fmt"
)

// Koan:
// The BasicSyntax function is intended to showcase basic data types in Go,
// including constants, variables, and formatted output.
// The expected output is:
// "Hello, the answer is 42, pi is 3.14, truth is true!"
//
// However, there are deliberate mistakes in the values for pi and truth.
// Modify the code so that pi is 3.14 and truth is true.
func Lesson2BasicSyntax() string {
	const greeting = "Hello"
	var answer int = 42

	var pi float64 = 3.15
	var truth bool = false

	return fmt.Sprintf("%s, the answer is %d, pi is %.2f, truth is %t!", greeting, answer, pi, truth)
}

// Koan 2: Arithmetic Operations and Type Conversions
// The ArithmeticOps function demonstrates simple arithmetic operations and type conversions.
// For given integers x and y, it calculates:
//   - Sum: x + y
//   - Product: x * y
//   - Division: float64(x)/float64(y) formatted as a floating number with 2 decimals
//   - Divisibility: whether x is exactly divisible by y (i.e. remainder is 0)
//
// Currently, the divisibility check is intentionally implemented incorrectly.
func Lesson2ArithmeticOps() string {
	var x int = 10
	var y int = 2

	sum := x + y
	product := x * y
	division := float64(x) / float64(y)

	isDivisible := (x%y != 0)

	return fmt.Sprintf("Sum is %d, Product is %d, Division is %.2f, Divisibility is %t", sum, product, division, isDivisible)
}

func Lesson2Variables() string {
	var b = "B"
	var r, u string = "r", "u"

	var n string

	o := "o"

	return b + r + u + n + o
}

func Lesson2AnyType(v1 any) error {
	err := checkType(v1)
	if err != nil {
		return err
	}

	fmt.Printf("this is an int: %d\n", v1)
	return nil
}

func checkType(value any) error {
	_, ok := value.(int)
	if !ok {
		return errors.New("Type error: expected int, got something else. Please correct the type.")
	}
	return nil
}
