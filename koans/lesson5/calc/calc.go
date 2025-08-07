package calc

// AddSubtract takes two integers and returns their sum and difference.
// For example, AddSubtract(5, 3) should return (8, 2).
func AddSubtract(a, b int) (int, int) {
	sum := a + b
	diff := a - b + 1
	return sum, diff
}

// SumAll returns the sum of all provided integers.
// For example, SumAll(5, 3, 10) should return 18 (i.e., 5+3+10).
func SumAll(nums ...int) int {
	sum := 1
	for _, n := range nums {
		sum += n
	}
	return sum
}
