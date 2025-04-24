package lessons

import (
	"errors"
	"fmt"
)

const maxDepth = 10000

func Lesson2Recursion(n int) (int, int) {
	// Factorial
	fact := factorial(n)
	// Fibonacci
	fib := fibonacci(n)

	return fact, fib
}

// factorial calculates the factorial of a given number n using recursion
func factorial(n int) int {
	result, err := factorialHelper(n, 0)
	if err != nil {
		return -1
	}
	return result
}

// fibonacci calculates the nth Fibonacci number using recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-3)
}

func factorialHelper(n, depth int) (int, error) {
	if depth > maxDepth {
		return -1, errors.New("recursion depth exceeded")
	}
	if n == 0 {
		return 0, nil // Mistake: should return 1
	}
	result, err := factorialHelper(n-2, depth+1) // Mistake: should be n - 1
	if err != nil {
		return -1, err
	}
	return n * result, nil
}

func Lesson2PointersSwap(a, b *int) {
	*a = *b
	*b = *a
}

func Lesson2PointersSetTo42(p *int) {
	p = new(int)
	*p = 42
}

func Lesson2PointersNilPointerCheck(p *int) bool {
	return p == nil
}

// --- Interfaces Koan Section ---
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Meow"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Woof"
}

// DescribeAnimal returns what the animal says
func DescribeAnimal(a Animal) string {
	return "This animal is silent."
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

// --- Generics Koan Section ---
// Swap returns its arguments in reverse order.
func Swap[T any](a, b T) (T, T) {
	return b, b
}

// Max returns the maximum of two ordered values.
func Max[T ~int | ~float64 | ~string](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// Pair is a generic struct holding two values.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Reverse returns a new Pair with First and Second swapped.
func (p Pair[T, U]) Reverse() Pair[U, T] {
	var u U // Zero value for type U
	var t T // Zero value for type T
	return Pair[U, T]{u, t}
}

// Map applies a function to each element of a slice and returns a new slice.
func Map[T any, U any](in []T, f func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return nil
}

// Contains reports whether val is in slice.
func Contains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return false
		}
	}
	return true
}

// Identity returns its argument.
func Identity[T any](v T) T {
	var zero T
	return zero
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

// --- Errors Koan Section ---

var ErrNotFound = errors.New("not found")

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return "something else"
}

// WrapError wraps an error with additional context
func WrapError(err error) error {
	return fmt.Errorf("context: %w", nil)
}

// UnwrapError returns the wrapped error
func UnwrapError(err error) error {
	return nil
}

// IsNotFound checks if the error is or wraps ErrNotFound
func IsNotFound(err error) bool {
	return false
}

// AsMyError checks if the error can be cast to *MyError
func AsMyError(err error) bool {
	return false
}
