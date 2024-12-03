package lessons

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
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

func compare(t *testing.T, expected, actual any, hint string) bool {
	if expected != actual {
		t.Errorf(`
		%sThe path to enlightenment is fraught with errors.
		Hint: %s%s
		%sExpected: %s%q
		%sBut found: %s%q
		%sReflect on this and try again.%s`, Green, Reset, hint, Blue, Reset, expected, Red, Reset, actual, Cyan, Reset)
		return false
	}
	t.Log("You have successfully passed this test, congratulations.")
	return true
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	original := os.Stdout
	os.Stdout = w

	f()

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = original

	return buf.String()
}
