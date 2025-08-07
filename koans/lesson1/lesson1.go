package lesson1

import "fmt"

// Koan:
// The output of the Hello function is currently incorrect.
// Fix the returned string so that when you run the program,
// it prints exactly "Hello, World!".
func Hello() string {
	// Buggy output: note the typo ("Helo" instead of "Hello")
	return "Helo, World!" // <-- Change this to "Hello, World!"
}

func init() {
	fmt.Println(Hello())
}
