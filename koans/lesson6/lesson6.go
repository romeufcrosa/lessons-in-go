package lesson6

import "fmt"

// Person demonstrates a custom struct with Name and Age.
type Person struct {
	Name string
	Age  int
}

// Greet returns a greeting message using the Person's fields.
// This method uses a value receiver which is fine because it does not modify the struct.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// Birthday is intended to increment the Person's Age by 1.
func (p Person) Birthday() {
	p.Age++
}

// CorrectBirthday demonstrates the proper use of a pointer receiver to modify the struct.
func (p *Person) CorrectBirthday() {
	p.Age++
}

func Lesson6PointersSwap(a, b *int) {
	*a = *b
	*b = *a
}

func Lesson6PointersSetTo42(p *int) {
	p = new(int)
	*p = 42
}

func Lesson6PointersNilPointerCheck(p *int) bool {
	return p == nil
}
