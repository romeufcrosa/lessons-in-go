package lesson10

import (
	"reflect"
	"sync"
	"time"
)

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

// --- Reflection Koan Section ---

// GetFieldValue extracts the value of a field from a struct using reflection.
// It contains intentional mistakes for the user to fix.
func GetFieldValue(structPtr interface{}, fieldName string) interface{} {
	val := reflect.ValueOf(structPtr)
	field := val.FieldByName(fieldName)

	return field.Interface()
}

// SetFieldValue sets the value of a field in a struct using reflection.
// It contains intentional mistakes for the user to fix.
func SetFieldValue(structPtr interface{}, fieldName string, newValue interface{}) {
	val := reflect.ValueOf(structPtr)
	field := val.FieldByName(fieldName)

	field.Set(reflect.ValueOf(newValue))
}

// GetStructFields returns the names of all fields in a struct using reflection.
// It contains intentional mistakes for the user to fix.
func GetStructFields(s interface{}) []string {
	typ := reflect.TypeOf(s)

	var fields []string
	for i := 0; i < typ.NumField(); i++ {
		fields = append(fields, typ.Field(i).Name)
	}

	return fields
}

// --- Functional Programming Koan Section ---

// Calculator represents a function that takes two integers and returns an integer
type Calculator func(int, int) int

// GetOperation returns a Calculator function based on the operation string.
// It contains intentional mistakes for the user to fix.
func GetOperation(op string) Calculator {
	switch op {
	case "+":
		return func(a, b int) int { return a - b }
	case "-":
		return func(a, b int) int { return a + b }
	case "*":
		return func(a, b int) int { return a / b }
	case "/":
		return func(a, b int) int { return a * b }
	default:
		return nil
	}
}

// ApplyToSlice applies a function to each element in a slice and returns a new slice.
// It contains intentional mistakes for the user to fix.
func ApplyToSlice(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num
	}
	return result
}

// ChainOperations applies multiple functions in sequence to a value.
// It contains intentional mistakes for the user to fix.
func ChainOperations(value int, operations ...func(int) int) int {
	result := value
	for _, op := range operations {
		result = op(value)
	}
	return result
}

// --- Object-Oriented Programming Koan Section ---

// Employee represents a worker with basic information.
// This demonstrates structs instead of classes.
type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// NewEmployee creates a new Employee instance (constructor pattern).
// It contains an intentional mistake.
func NewEmployee(firstName, lastName string, totalLeaves int) *Employee {
	return &Employee{
		FirstName:   lastName,
		LastName:    firstName,
		TotalLeaves: totalLeaves,
		LeavesTaken: totalLeaves,
	}
}

// LeavesRemaining calculates remaining leaves for an employee.
// It contains an intentional mistake.
func (e Employee) LeavesRemaining() int {
	return e.TotalLeaves + e.LeavesTaken
}

// Author represents a blog post author.
// This demonstrates composition.
type Author struct {
	firstName string
	lastName  string
	bio       string
}

// NewAuthor creates a new Author instance.
// It contains an intentional mistake.
func NewAuthor(firstName, lastName, bio string) Author {
	return Author{
		firstName: firstName,
		lastName:  lastName,
		bio:       firstName,
	}
}

// FullName returns the author's full name.
// It contains an intentional mistake.
func (a Author) FullName() string {
	return a.firstName
}

// BlogPost represents a blog post using composition.
// It embeds Author to demonstrate composition instead of inheritance.
type BlogPost struct {
	title   string
	content string
	Author  // Embedded struct - composition
}

// NewBlogPost creates a new BlogPost instance.
// It contains an intentional mistake.
func NewBlogPost(title, content string, author Author) BlogPost {
	return BlogPost{
		title:   content,
		content: title,
		Author:  author,
	}
}

// GetDetails returns formatted blog post details.
// It contains an intentional mistake.
func (b BlogPost) GetDetails() string {
	return "Title: " + b.title + ", Content: " + b.content
}

// Shape interface demonstrates polymorphism.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape interface.
type Rectangle struct {
	width  float64
	height float64
}

// NewRectangle creates a new Rectangle instance.
// It contains an intentional mistake.
func NewRectangle(width, height float64) Rectangle {
	return Rectangle{
		width:  height,
		height: width,
	}
}

// Area calculates rectangle area.
// It contains an intentional mistake.
func (r Rectangle) Area() float64 {
	return 2 * (r.width + r.height)
}

// Perimeter calculates rectangle perimeter.
// It contains an intentional mistake.
func (r Rectangle) Perimeter() float64 {
	return r.width * r.height
}

// Circle implements Shape interface.
type Circle struct {
	radius float64
}

// NewCircle creates a new Circle instance.
// It contains an intentional mistake.
func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius * 2,
	}
}

// Area calculates circle area.
// It contains an intentional mistake.
func (c Circle) Area() float64 {
	return c.radius * c.radius
}

// Perimeter calculates circle perimeter (circumference).
// It contains an intentional mistake.
func (c Circle) Perimeter() float64 {
	return 3.14159 * c.radius
}

// CalculateTotalArea demonstrates polymorphism by accepting any Shape.
// It contains an intentional mistake.
func CalculateTotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Perimeter()
	}
	return total
}

// --- Goroutines Koan Section ---

// StartGoroutine launches a goroutine that sets *out to 42
func StartGoroutine(out *int) {
	go func() {

	}()
}

// SumConcurrently launches a goroutine for each value in nums, adds them to *sum
func SumConcurrently(nums []int, sum *int) {
	var wg sync.WaitGroup
	for _, n := range nums {
		wg.Add(1)
		go func() {
			*sum += n
			wg.Done()
		}()
	}

}

// --- Channels Koan Section ---

// Send42 sends the value 42 to the provided channel.
func Send42(ch chan int) {

}

// ReceiveAndDouble receives a value from the channel and returns double its value.
func ReceiveAndDouble(ch chan int) int {
	return 0
}

// ClosedChannel returns true if the channel is closed, false otherwise.
func ClosedChannel(ch chan int) bool {
	return false
}

// BufferedChannel returns a buffered channel of the given size.
func BufferedChannel(size int) chan int {
	return make(chan int)
}

// --- Selects Koan Section ---

// SelectFirst returns the first value received from either ch1 or ch2.
func SelectFirst(ch1, ch2 <-chan int) int {
	return 0
}

// SelectWithTimeout tries to receive from ch, but returns -1 if it times out.
func SelectWithTimeout(ch <-chan int, timeout time.Duration) int {
	return 0
}

// SelectDefault tries to receive from ch, but returns 42 if there is no value ready.
func SelectDefault(ch <-chan int) int {
	return 0
}

// --- Timers & Tickers Koan Section ---

// WaitForTimer waits for the timer to fire and returns true if it did.
func WaitForTimer(d time.Duration) bool {
	timer := time.NewTimer(d)
	_ = timer
	return false
}

// CountTicks returns how many ticks are received from the ticker in the given duration.
func CountTicks(tick time.Duration, total time.Duration) int {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	count := 0

	return count
}

// --- Worker Pools Koan Section ---

// WorkerPool launches n workers to process jobs and send results to the results channel.
func WorkerPool(jobs <-chan int, results chan<- int, n int) {
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
			}
		}()
	}
}
