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
	compare(t, "Hello Bruno, welcome to Lessons in Go", output)
}

func compare(t *testing.T, expected, actual interface{}, _ ...any) bool {
	if expected != actual {
		t.Errorf(`
		%sThe path to enlightenment is fraught with errors.
		%sExpected: %s%q
		%sBut found: %s%q
		%sReflect on this and try again.%s`, Green, Blue, Reset, expected, Red, Reset, actual, Cyan, Reset)
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
