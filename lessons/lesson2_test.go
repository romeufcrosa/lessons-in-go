package lessons

import (
	"errors"
	"fmt"
	"testing"
)

// func TestLesson2Recursion(t *testing.T) {
// 	fact, fib := Lesson2Recursion(5)
// 	compare(t, 120, fact, "Beware of infinite recursion")
// 	compare(t, 5, fib, "I sense your fibonacci calculation is off")

// 	fact, fib = Lesson2Recursion(0)
// 	compare(t, 1, fact, "almost there, just that edge case missing")
// 	compare(t, 0, fib, "almost there, just that edge case missing")
// }

func TestLesson2Pointers(t *testing.T) {
	// Test pointer swap
	a, b := 1, 2
	Lesson2PointersSwap(&a, &b)
	compare(t, 2, a, "Expected a=2 after swap")
	compare(t, 1, b, "Expected b=1 after swap")

	// Test SetTo42
	x := 10
	Lesson2PointersSetTo42(&x)
	compare(t, 42, x, "Expected x=42 after SetTo42")

	// Test nil pointer check
	var p *int
	compare(t, true, Lesson2PointersNilPointerCheck(p), "Expected true for nil pointer")

	x = 5
	compare(t, false, Lesson2PointersNilPointerCheck(&x), "Expected false for non-nil pointer")
}

func TestLesson2Interfaces(t *testing.T) {
	var a Animal

	a = Dog{}
	compare(t, "Woof", a.Speak(), "Dogs should say Woof")

	a = Cat{}
	compare(t, "Meow", a.Speak(), "Cats should say Meow")

	compare(t, "Woof", DescribeAnimal(Dog{}), "DescribeAnimal should use the animal's Speak method")
	compare(t, "Meow", DescribeAnimal(Cat{}), "DescribeAnimal should use the animal's Speak method")
}

func TestLesson2Enums(t *testing.T) {
	compare(t, "Sun", Sunday.String(), "Sunday should be 'Sun'")
	compare(t, "Mon", Monday.String(), "Monday should be 'Mon'")
	compare(t, "Sat", Saturday.String(), "Saturday should be 'Sat'")

	compare(t, Monday, NextDay(Sunday), "NextDay(Sunday) should be Monday")
	compare(t, Sunday, NextDay(Saturday), "NextDay(Saturday) should be Sunday")
}

func TestLesson2Generics(t *testing.T) {
	// Test Swap
	a, b := Swap(1, 2)
	compare(t, 2, a, "Swap should swap values")
	compare(t, 1, b, "Swap should swap values")

	// Test Max
	compare(t, 3, Max(2, 3), "Max should return the maximum value")
	compare(t, 2.5, Max(2.5, 1.5), "Max should return the maximum value")

	// Test Pair and Swap method
	p := Pair[int, string]{First: 1, Second: "hello"}
	swapped := p.Reverse()
	compare(t, "hello", swapped.First, "Pair.Swap should swap elements")
	compare(t, 1, swapped.Second, "Pair.Swap should swap elements")

	// Test Map
	double := func(x int) int { return x * 2 }
	result := Map([]int{1, 2, 3}, double)
	compareDeep(t, []int{2, 4, 6}, result, "Map should apply function to all elements")

	// Test Contains
	compare(t, true, Contains([]int{1, 2, 3}, 2), "Contains should return true if value is present")
	compare(t, false, Contains([]int{1, 2, 3}, 4), "Contains should return false if value is not present")
}

func TestLesson2Iterators(t *testing.T) {
	// Test IntRange
	var got []int
	for v := range IntRange(2, 5) {
		got = append(got, v)
	}
	//compareDeep(t, []int{2, 3, 4}, got, "IntRange should yield numbers from start to end-1")

	// Test SliceIterator
	iter := SliceIterator([]string{"a", "b", "c"})
	var s []string
	for {
		val, ok := iter()
		if !ok {
			break
		}
		s = append(s, val)
	}
	compareDeep(t, []string{"a", "b", "c"}, s, "SliceIterator should yield all elements in order")

	// Test MapKeysIterator
	m := map[string]int{"x": 1, "y": 2}
	keyIter := MapKeysIterator(m)
	keys := make(map[string]bool)
	for {
		k, ok := keyIter()
		if !ok {
			break
		}
		keys[k] = true
	}
	for k := range m {
		compare(t, true, keys[k], fmt.Sprintf("MapKeysIterator missing key: %v", k))
	}
}

func TestLesson2Errors(t *testing.T) {
	// Test WrapError and UnwrapError
	wrapped := WrapError(ErrNotFound)
	compare(t, true, errors.Is(wrapped, ErrNotFound), "WrapError should wrap ErrNotFound")
	compare(t, ErrNotFound, UnwrapError(wrapped), "UnwrapError should return the original error")

	// Test IsNotFound
	compare(t, true, IsNotFound(wrapped), "IsNotFound should return true for wrapped ErrNotFound")

	// Test AsMyError
	myErr := &MyError{Msg: "custom"}
	wrappedMyErr := WrapError(myErr)
	compare(t, true, AsMyError(wrappedMyErr), "AsMyError should return true for wrapped MyError")
}
