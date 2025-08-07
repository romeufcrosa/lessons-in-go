package lesson7

import (
	"testing"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

func TestLesson7Interfaces(t *testing.T) {
	var a Animal

	a = Dog{}
	testutils.Compare(t, "Woof", a.Speak(), "Dogs should say Woof")

	a = Cat{}
	testutils.Compare(t, "Meow", a.Speak(), "Cats should say Meow")

	testutils.Compare(t, "Woof", Lesson7DescribeAnimal(Dog{}), "DescribeAnimal should use the animal's Speak method")
	testutils.Compare(t, "Meow", Lesson7DescribeAnimal(Cat{}), "DescribeAnimal should use the animal's Speak method")
}

// TestPersonSpeak verifies that Person.Speak returns the correct string.
// Expected for Person{Name: "Alice"}: "I am Alice"
func TestLesson7PersonSpeak(t *testing.T) {
	p := Person{Name: "Alice"}
	expected := "I am Alice"
	actual := p.Speak()
	hint := "Person.Speak should return 'I am <Name>'. Ensure that you include 'am' after 'I'."
	testutils.Compare(t, expected, actual, hint)
}

// TestRobotSpeak verifies that Robot.Speak returns the correct string.
// Expected for Robot{ID: 42}: "Beep boop, I am robot 42"
func TestLesson7RobotSpeak(t *testing.T) {
	r := Robot{ID: 42}
	expected := "Beep boop, I am robot 42"
	actual := r.Speak()
	hint := "Robot.Speak should return 'Beep boop, I am robot <ID>'. Fix the wording in the output."
	testutils.Compare(t, expected, actual, hint)
}

// TestIdentifySpeaker confirms that IdentifySpeaker properly distinguishes types.
func TestLesson7IdentifySpeaker(t *testing.T) {
	p := Person{Name: "Alice"}
	r := Robot{ID: 42}
	expectedP := "Person: Alice"
	expectedR := "Robot: 42"
	actualP := Lesson7IdentifySpeaker(p)
	actualR := Lesson7IdentifySpeaker(r)
	hintP := "IdentifySpeaker should recognize a Person and return 'Person: <Name>'."
	hintR := "IdentifySpeaker should recognize a Robot and return 'Robot: <ID>'."
	testutils.Compare(t, expectedP, actualP, hintP)
	testutils.Compare(t, expectedR, actualR, hintR)
}

// TestEmployeeSpeak verifies that Employee.Speak returns the correctly composed string.
// For an Employee with Name "Bob" and Position "Engineer", expected output:
// "I am Bob and I work as Engineer"
func TestLesson7EmployeeSpeak(t *testing.T) {
	e := Employee{
		Person:   Person{Name: "Bob"},
		Position: "Engineer",
	}
	expected := "I am Bob and I work as Engineer"
	actual := e.Speak()
	hint := "Employee.Speak should utilize the embedded Person and include the phrase 'work as'. Ensure you format it as 'I am <Name> and I work as <Position>'."
	testutils.Compare(t, expected, actual, hint)
}
