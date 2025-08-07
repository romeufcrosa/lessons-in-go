package lesson7

import "fmt"

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
func Lesson7DescribeAnimal(a Animal) string {
	return "This animal is silent."
}

// Speaker is an interface with a Speak method.
type Speaker interface {
	Speak() string
}

// Person is a concrete type that implements Speaker.
type Person struct {
	Name string
}

// Intentional bug: Person.Speak should include "am".
// Expected output for Person{Name: "Alice"}: "I am Alice"
func (p Person) Speak() string {
	return "I " + p.Name // <-- Bug: missing "am"
}

// Robot is another type that implements Speaker.
type Robot struct {
	ID int
}

// Intentional bug: Robot.Speak should return "Beep boop, I am robot <ID>".
// Currently it returns "Beep beep, I am robot <ID>".
func (r Robot) Speak() string {
	return fmt.Sprintf("Beep beep, I am robot %d", r.ID) // <-- Bug: "Beep beep" instead of "Beep boop"
}

// IdentifySpeaker uses a type switch to return specific identity information
// based on the concrete type behind Speaker.
func Lesson7IdentifySpeaker(s Speaker) string {
	var identity string
	switch v := s.(type) {
	case Person:
		identity = fmt.Sprintf("Person: %s", v.Name)
	case Robot:
		identity = fmt.Sprintf("Robot: %d", v.ID)
	default:
		identity = "Unknown"
	}
	return identity
}

// Employee demonstrates composition by embedding a Person.
// It should override the Speak method to incorporate the employee's position.
// Expected output for an Employee with Name "Bob" and Position "Engineer":
// "I am Bob and I work as Engineer"
//
// Intentional bug: The formatted string omits the phrase "work as".
func (e Employee) Speak() string {
	return fmt.Sprintf("I am %s and I %s", e.Name, e.Position) // <-- Bug: missing "work as"
}

// Employee composes a Person and additional fields.
type Employee struct {
	Person   // Embedding Person, so Employee reuses Person's fields and methods.
	Position string
}
