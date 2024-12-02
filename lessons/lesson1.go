package lessons

import (
	"errors"
	"fmt"
)

func Lesson1PrintFunc(_ string) {
	fmt.Printf("Hello %s, welcome to Lessons in Go", "")
}

func Lesson1VarTypes(v1 any) error {
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

func checkType(value interface{}) error {
	_, ok := value.(int)
	// Change int to the expected type
	if !ok {
		return errors.New("Type error: expected int, got something else. Please correct the type.")
	}
	return nil
}
