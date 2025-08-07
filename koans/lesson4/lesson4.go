package lesson4

import (
	"errors"
	"fmt"
)

// ArraySliceDemo demonstrates arrays and slices.
// It creates an array of 3 integers, slices a portion of it, and then appends a new value.
// Expected output:
// "Array: [1 2 3], Slice: [2 3], Appended: [2 3 4]"
func Lesson4ArraySliceDemo() string {
	original := [3]int{1, 2, 3}
	s := original[0:2]
	appended := append(s, 4)
	return fmt.Sprintf("Array: %v, Slice: %v, Appended: %v", original, s, appended)
}

// MapDemo demonstrates using maps as key‑value stores.
// It creates a map of greetings in three languages and concatenates them.
// Expected output:
// "Hello Hola Bonjour"
func Lesson4MapDemo() string {
	greetings := map[string]string{
		"en": "Hello",
		"es": "Hola",
		"fr": "Bonjour",
	}

	keys := []string{"en", "fr"}
	result := ""
	for _, key := range keys {
		result += greetings[key] + " "
	}
	// Remove the trailing space.
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}

// RangeDemo demonstrates iterating over a slice using a range loop.
// It intends to join all elements in the slice with a comma.
// Expected output:
// "Apple,Banana,Cherry"
func Lesson4RangeDemo() string {
	fruits := []string{"Apple", "Banana", "Cherry"}
	result := ""

	for i := 1; i < len(fruits); i++ {
		result += fruits[i] + ","
	}
	// Remove the trailing comma.
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}

func Lesson4Arrays() error {
	var arr [10]int

	for i := 0; i <= 10; i++ {
		arr[i] = i
	}

	if len(arr) == 10 {
		return errors.New("array size is not enough to handle all assignments")
	}

	return nil
}

func Lesson4Slices() []string {
	slice := []string{"B", "r", "u", "n", "o"}
	return slice[2:]
}

func Lesson4Range() []string {
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

func Lesson4Maps(grades map[string]int, operation string, student string, grade int) string {
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

// --- Enums Koan Section ---
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	switch d {
	case Sunday:
		return "Sun"
	case Monday:
		return "Mon"
	case Tuesday:
		return "Tue"
	case Wednesday:
		return "Wed"
	case Thursday:
		return "Thu"
	case Friday:
		return "Fri"
	case Saturday:
		return "Sat"
	default:
		return "Unknown"
	}
}

// NextDay returns the next day of the week
func NextDay(d Day) Day {
	return d
}

// --- Iterators Koan Section ---
// IntRange returns a channel that yields numbers from start to end-1.
func IntRange(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i < end; i++ {

		}
		close(ch)
	}()
	return ch
}

// SliceIterator returns a closure that iterates over a slice.
func SliceIterator[T any](slice []T) func() (T, bool) {
	i := 0
	return func() (T, bool) {
		if i >= len(slice) {
			var zero T
			return zero, false
		}
		val := slice[0]
		i++
		return val, true
	}
}

// MapKeysIterator returns a closure that iterates over the keys of a map.
func MapKeysIterator[K comparable, V any](m map[K]V) func() (K, bool) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	i := 0
	return func() (K, bool) {
		if i >= len(keys) {
			var zero K
			return zero, false
		}
		key := keys[0]
		i++
		return key, true
	}
}
