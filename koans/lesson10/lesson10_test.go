package lesson10

import (
	"fmt"
	"testing"
	"time"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// Test struct for reflection tests
type Person struct {
	Name    string
	Age     int
	private string
}

// TestGetFieldValue tests reflection field access
func TestGetFieldValue(t *testing.T) {
	person := &Person{Name: "Alice", Age: 30}

	name := GetFieldValue(person, "Name")
	testutils.Compare(t, "Alice", name, "GetFieldValue should extract the Name field correctly")

	age := GetFieldValue(person, "Age")
	testutils.Compare(t, 30, age, "GetFieldValue should extract the Age field correctly")
}

// TestSetFieldValue tests reflection field modification
func TestSetFieldValue(t *testing.T) {
	person := &Person{Name: "Bob", Age: 25}

	SetFieldValue(person, "Name", "Charlie")
	testutils.Compare(t, "Charlie", person.Name, "SetFieldValue should update the Name field")

	SetFieldValue(person, "Age", 35)
	testutils.Compare(t, 35, person.Age, "SetFieldValue should update the Age field")
}

// TestGetStructFields tests reflection field enumeration
func TestGetStructFields(t *testing.T) {
	person := Person{Name: "Dave", Age: 40}

	fields := GetStructFields(person)
	expected := []string{"Name", "Age"} // Should only include exported fields
	testutils.Compare(t, expected, fields, "GetStructFields should return only exported field names")

	// Test with pointer
	fields = GetStructFields(&person)
	testutils.Compare(t, expected, fields, "GetStructFields should handle pointers correctly")
}

func TestLesson2Generics(t *testing.T) {
	// Test Swap
	a, b := Swap(1, 2)
	testutils.Compare(t, 2, a, "Swap should swap values")
	testutils.Compare(t, 1, b, "Swap should swap values")

	// Test Max
	testutils.Compare(t, 3, Max(2, 3), "Max should return the maximum value")
	testutils.Compare(t, 2.5, Max(2.5, 1.5), "Max should return the maximum value")

	// Test Pair and Swap method
	p := Pair[int, string]{First: 1, Second: "hello"}
	swapped := p.Reverse()
	testutils.Compare(t, "hello", swapped.First, "Pair.Swap should swap elements")
	testutils.Compare(t, 1, swapped.Second, "Pair.Swap should swap elements")

	// Test Map
	double := func(x int) int { return x * 2 }
	result := Map([]int{1, 2, 3}, double)
	testutils.CompareDeep(t, []int{2, 4, 6}, result, "Map should apply function to all elements")

	// Test Contains
	testutils.Compare(t, true, Contains([]int{1, 2, 3}, 2), "Contains should return true if value is present")
	testutils.Compare(t, false, Contains([]int{1, 2, 3}, 4), "Contains should return false if value is not present")
}

// TestGetOperation tests first-class functions (functions as values)
func TestGetOperation(t *testing.T) {
	add := GetOperation("+")
	result := add(5, 3)
	testutils.Compare(t, 8, result, "Addition operation should return 8")

	subtract := GetOperation("-")
	result = subtract(10, 4)
	testutils.Compare(t, 6, result, "Subtraction operation should return 6")

	multiply := GetOperation("*")
	result = multiply(3, 4)
	testutils.Compare(t, 12, result, "Multiplication operation should return 12")

	divide := GetOperation("/")
	result = divide(15, 3)
	testutils.Compare(t, 5, result, "Division operation should return 5")
}

// TestApplyToSlice tests higher-order functions (functions that take functions as parameters)
func TestApplyToSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Test with square function
	square := func(x int) int { return x * x }
	squared := ApplyToSlice(numbers, square)
	expected := []int{1, 4, 9, 16, 25}
	testutils.Compare(t, expected, squared, "ApplyToSlice should square each number")

	// Test with double function
	double := func(x int) int { return x * 2 }
	doubled := ApplyToSlice(numbers, double)
	expected = []int{2, 4, 6, 8, 10}
	testutils.Compare(t, expected, doubled, "ApplyToSlice should double each number")
}

// TestChainOperations tests function composition and chaining
func TestChainOperations(t *testing.T) {
	add5 := func(x int) int { return x + 5 }
	multiply2 := func(x int) int { return x * 2 }
	subtract3 := func(x int) int { return x - 3 }

	// Start with 10, add 5 (=15), multiply by 2 (=30), subtract 3 (=27)
	result := ChainOperations(10, add5, multiply2, subtract3)
	testutils.Compare(t, 27, result, "ChainOperations should apply functions in sequence")

	// Test with single operation
	result = ChainOperations(5, multiply2)
	testutils.Compare(t, 10, result, "ChainOperations should work with single operation")
}

// --- Object-Oriented Programming Tests ---

// TestEmployeeStructs tests structs instead of classes pattern
func TestEmployeeStructs(t *testing.T) {
	// Test constructor pattern
	emp := NewEmployee("John", "Doe", 25)
	testutils.Compare(t, "John", emp.FirstName, "Employee FirstName should be set correctly")
	testutils.Compare(t, "Doe", emp.LastName, "Employee LastName should be set correctly")
	testutils.Compare(t, 25, emp.TotalLeaves, "Employee TotalLeaves should be set correctly")
	testutils.Compare(t, 0, emp.LeavesTaken, "Employee LeavesTaken should start at 0")

	// Test method
	emp.LeavesTaken = 5
	remaining := emp.LeavesRemaining()
	testutils.Compare(t, 20, remaining, "LeavesRemaining should calculate correctly")
}

// TestComposition tests composition instead of inheritance pattern
func TestComposition(t *testing.T) {
	// Test Author creation
	author := NewAuthor("Jane", "Smith", "Tech blogger and Go enthusiast")
	testutils.Compare(t, "Jane Smith", author.FullName(), "Author FullName should return full name")

	// Test BlogPost composition
	post := NewBlogPost("Go OOP Patterns", "Content about Go's approach to OOP", author)
	testutils.Compare(t, "Go OOP Patterns", post.title, "BlogPost title should be set correctly")
	testutils.Compare(t, "Content about Go's approach to OOP", post.content, "BlogPost content should be set correctly")

	// Test embedded functionality
	testutils.Compare(t, "Jane Smith", post.FullName(), "BlogPost should have access to Author methods through embedding")

	// Test GetDetails method
	details := post.GetDetails()
	expectedDetails := "Title: Go OOP Patterns, Content: Content about Go's approach to OOP, Author: Jane Smith"
	testutils.Compare(t, expectedDetails, details, "GetDetails should include author information")
}

// TestPolymorphism tests polymorphism with interfaces
func TestPolymorphism(t *testing.T) {
	// Test Rectangle
	rect := NewRectangle(4.0, 3.0)
	testutils.Compare(t, 12.0, rect.Area(), "Rectangle area should be width * height")
	testutils.Compare(t, 14.0, rect.Perimeter(), "Rectangle perimeter should be 2 * (width + height)")

	// Test Circle
	circle := NewCircle(5.0)
	expectedArea := 3.14159 * 5.0 * 5.0 // π * r²
	testutils.Compare(t, expectedArea, circle.Area(), "Circle area should be π * r²")

	expectedPerimeter := 2 * 3.14159 * 5.0 // 2π * r
	testutils.Compare(t, expectedPerimeter, circle.Perimeter(), "Circle perimeter should be 2π * r")

	// Test polymorphism - interface can hold different concrete types
	var shapes []Shape
	shapes = append(shapes, rect)
	shapes = append(shapes, circle)

	// Test CalculateTotalArea function
	totalArea := CalculateTotalArea(shapes)
	expectedTotal := rect.Area() + circle.Area()
	testutils.Compare(t, expectedTotal, totalArea, "CalculateTotalArea should sum areas of all shapes")
}

func TestLesson10Goroutines(t *testing.T) {
	var out int
	StartGoroutine(&out)
	// Give goroutine time to run (not ideal, but for Koans)
	time.Sleep(50 * time.Millisecond)
	testutils.Compare(t, 42, out, "StartGoroutine should set *out to 42")

	nums := []int{1, 2, 3, 4}
	var sum int
	SumConcurrently(nums, &sum)
	time.Sleep(50 * time.Millisecond)
	testutils.Compare(t, 10, sum, "SumConcurrently should sum all numbers concurrently")
}

func TestLesson10Channels(t *testing.T) {
	ch := make(chan int)
	go Send42(ch)
	select {
	case val := <-ch:
		testutils.Compare(t, 42, val, "Send42 should send 42 to the channel")
	case <-time.After(500 * time.Millisecond):
		t.Errorf("Timeout: Did you forget to send a value to the channel in Send42?")
	}

	ch2 := make(chan int)
	go func() { ch2 <- 21 }()
	result := ReceiveAndDouble(ch2)
	testutils.Compare(t, 42, result, "ReceiveAndDouble should double the received value")

	ch3 := make(chan int)
	close(ch3)
	testutils.Compare(t, true, ClosedChannel(ch3), "ClosedChannel should return true for closed channel")
	ch4 := make(chan int)
	testutils.Compare(t, false, ClosedChannel(ch4), "ClosedChannel should return false for open channel")

	bufCh := BufferedChannel(3)
	capacity := cap(bufCh)
	testutils.Compare(t, 3, capacity, "BufferedChannel should return a channel with the correct buffer size")
}

func TestLesson10Selects(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { ch2 <- 7 }()
	testutils.Compare(t, 7, SelectFirst(ch1, ch2), "SelectFirst should return the first value received from either channel")

	ch := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 5
	}()
	testutils.Compare(t, -1, SelectWithTimeout(make(chan int), 50*time.Millisecond), "SelectWithTimeout should return -1 on timeout")
	testutils.Compare(t, 5, SelectWithTimeout(ch, 200*time.Millisecond), "SelectWithTimeout should return value if received before timeout")

	ch3 := make(chan int)
	testutils.Compare(t, 42, SelectDefault(ch3), "SelectDefault should return 42 if no value is ready")
	go func() { ch3 <- 99 }()
	time.Sleep(10 * time.Millisecond)
	testutils.Compare(t, 99, SelectDefault(ch3), "SelectDefault should return value if ready")
}

func TestLesson10TimersAndTickers(t *testing.T) {
	start := time.Now()
	ok := WaitForTimer(50 * time.Millisecond)
	elapsed := time.Since(start)
	if !ok || elapsed < 50*time.Millisecond {
		t.Errorf("WaitForTimer should wait for the timer and return true")
	}

	count := CountTicks(10*time.Millisecond, 55*time.Millisecond)
	// Should get about 5-6 ticks depending on timing
	testutils.Compare(t, true, count >= 4 && count <= 7, fmt.Sprintf("CountTicks should count ticks from the ticker, got %d", count))
}

func TestLesson10WorkerPools(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	WorkerPool(jobs, results, 3)

	var out []int
	for range 5 {
		out = append(out, <-results)
	}
	testutils.CompareDeep(t, []int{1, 4, 9, 16, 25}, out, "WorkerPool should square each job and send to results")
}
