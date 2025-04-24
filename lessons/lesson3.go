package lessons

import (
	"sync"
	"time"
)

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
