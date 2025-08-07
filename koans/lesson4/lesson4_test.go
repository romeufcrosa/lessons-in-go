package lesson4

import (
	"fmt"
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// TestArraySliceDemo verifies that ArraySliceDemo returns the expected output.
func TestLesson4ArraySliceDemo(t *testing.T) {
	expected := "Array: [1 2 3], Slice: [2 3], Appended: [2 3 4]"
	actual := Lesson4ArraySliceDemo()
	hint := "Adjust the slicing to select elements from index 1 to 3 (i.e., use original[1:3]) so that the slice becomes [2 3] and appending 4 yields [2 3 4]."
	testutils.Compare(t, expected, actual, hint)
}

// TestMapDemo verifies that MapDemo returns the expected greeting string.
func TestLesson4MapDemo(t *testing.T) {
	expected := "Hello Hola Bonjour"
	actual := Lesson4MapDemo()
	hint := "Ensure that the keys slice includes all language codes (i.e., [\"en\", \"es\", \"fr\"]) so that the concatenated greetings include 'Hola'."
	testutils.Compare(t, expected, actual, hint)
}

// TestRangeDemo verifies that RangeDemo returns the expected comma‑separated string.
func TestLesson4RangeDemo(t *testing.T) {
	expected := "Apple,Banana,Cherry"
	actual := Lesson4RangeDemo()
	hint := "Iterate over the entire slice (for example, by using a proper range loop) so that all elements, including the first one, are included."
	testutils.Compare(t, expected, actual, hint)
}

func TestLesson4Array(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			testutils.Compare(t, nil, r, "arrays have limits")
		}
	}()
	err := Lesson4Arrays()
	testutils.Compare(t, nil, err, "you have found an alternative path, but try modifying the bounds first")
}

func TestLesson4Slices(t *testing.T) {
	output := Lesson4Slices()
	testutils.Compare(t, "uno", output, "a slice of a slice is still a slice")
}

func TestLesson4Range(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			testutils.Compare(t, []string{"Rosa", "Bruno", "Igor"}, r, "declaration is not initialization")
		}
	}()
	output := Lesson4Range()
	testutils.Compare(t, []string{"Rosa", "Bruno", "Igor"}, output, "range is useful, but not in this scenario. do you know why?")
}

func TestLesson4Maps(t *testing.T) {
	grades := map[string]int{
		"Bruno": 15,
		"Igor":  12,
	}

	// Test adding a new student
	result := Lesson4Maps(grades, "add", "Romeu", 10)
	testutils.Compare(t, "Added Romeu with grade 10", result, "")

	// Test updating an existing student
	result = Lesson4Maps(grades, "update", "Bruno", 16)
	testutils.Compare(t, "Updated Bruno to grade 16", result, "")

	// Test updating a non-existing student (should fail)
	result = Lesson4Maps(grades, "update", "Margarida", 18)
	testutils.Compare(t, "Student Margarida not found", result, "I don't remember creating Margarida")

	// Test getting a non-existing student's grade (should fail)
	result = Lesson4Maps(grades, "get", "Margarida", 0)
	testutils.Compare(t, "Student Margarida not found", result, "still don't think Margarida actually exists")

}

func TestLesson2Enums(t *testing.T) {
	testutils.Compare(t, "Sun", Sunday.String(), "Sunday should be 'Sun'")
	testutils.Compare(t, "Mon", Monday.String(), "Monday should be 'Mon'")
	testutils.Compare(t, "Sat", Saturday.String(), "Saturday should be 'Sat'")

	testutils.Compare(t, Monday, NextDay(Sunday), "NextDay(Sunday) should be Monday")
	testutils.Compare(t, Sunday, NextDay(Saturday), "NextDay(Saturday) should be Sunday")
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
	testutils.CompareDeep(t, []string{"a", "b", "c"}, s, "SliceIterator should yield all elements in order")

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
		testutils.Compare(t, true, keys[k], fmt.Sprintf("MapKeysIterator missing key: %v", k))
	}
}
