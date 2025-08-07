package testutils

import (
	"bytes"
	"io"
	"log"
	"os"
	"reflect"
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

func Compare(t *testing.T, expected, actual any, hint string) bool {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(`
        %sThe path to enlightenment is fraught with errors.
        Hint: %s%s
        %sExpected: %s%v
        %sBut found: %s%v
        %sReflect on this and try again.%s`, Green, Reset, hint, Blue, Reset, expected, Red, Reset, actual, Cyan, Reset)
		return false
	}
	t.Log("You have successfully passed this test, congratulations.")
	return true
}

func CompareDeep(t *testing.T, expected, actual interface{}, msg string) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(`
		%sThe path to enlightenment is fraught with errors.
		Hint: %s%s
		%sExpected: %s%v
		%sBut found: %s%v
		%sReflect on this and try again.%s`, Green, Reset, msg, Blue, Reset, expected, Red, Reset, actual, Cyan, Reset)
		return
	}
	t.Log("You have successfully passed this test, congratulations.")
}

func PrintFailure(t *testing.T, expected, actual any, hint string) {
	t.Errorf(`
		%sThe path to enlightenment is fraught with errors.
		Hint: %s%s
		%sExpected: %s%v
		%sBut found: %s%v
		%sReflect on this and try again.%s`, Green, Reset, hint, Blue, Reset, expected, Red, Reset, actual, Cyan, Reset)
}

func CaptureOutput(f func()) string {
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
