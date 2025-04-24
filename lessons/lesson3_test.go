package lessons

import (
	"fmt"
	"testing"
	"time"
)

func TestLesson3Goroutines(t *testing.T) {
	var out int
	StartGoroutine(&out)
	// Give goroutine time to run (not ideal, but for Koans)
	time.Sleep(50 * time.Millisecond)
	compare(t, 42, out, "StartGoroutine should set *out to 42")

	nums := []int{1, 2, 3, 4}
	var sum int
	SumConcurrently(nums, &sum)
	time.Sleep(50 * time.Millisecond)
	compare(t, 10, sum, "SumConcurrently should sum all numbers concurrently")
}

func TestLesson3Channels(t *testing.T) {
	ch := make(chan int)
	go Send42(ch)
	select {
	case val := <-ch:
		compare(t, 42, val, "Send42 should send 42 to the channel")
	case <-time.After(500 * time.Millisecond):
		t.Errorf("Timeout: Did you forget to send a value to the channel in Send42?")
	}

	ch2 := make(chan int)
	go func() { ch2 <- 21 }()
	result := ReceiveAndDouble(ch2)
	compare(t, 42, result, "ReceiveAndDouble should double the received value")

	ch3 := make(chan int)
	close(ch3)
	compare(t, true, ClosedChannel(ch3), "ClosedChannel should return true for closed channel")
	ch4 := make(chan int)
	compare(t, false, ClosedChannel(ch4), "ClosedChannel should return false for open channel")

	bufCh := BufferedChannel(3)
	capacity := cap(bufCh)
	compare(t, 3, capacity, "BufferedChannel should return a channel with the correct buffer size")
}

func TestLesson3Selects(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { ch2 <- 7 }()
	compare(t, 7, SelectFirst(ch1, ch2), "SelectFirst should return the first value received from either channel")

	ch := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 5
	}()
	compare(t, -1, SelectWithTimeout(make(chan int), 50*time.Millisecond), "SelectWithTimeout should return -1 on timeout")
	compare(t, 5, SelectWithTimeout(ch, 200*time.Millisecond), "SelectWithTimeout should return value if received before timeout")

	ch3 := make(chan int)
	compare(t, 42, SelectDefault(ch3), "SelectDefault should return 42 if no value is ready")
	go func() { ch3 <- 99 }()
	time.Sleep(10 * time.Millisecond)
	compare(t, 99, SelectDefault(ch3), "SelectDefault should return value if ready")
}

func TestLesson3TimersAndTickers(t *testing.T) {
	start := time.Now()
	ok := WaitForTimer(50 * time.Millisecond)
	elapsed := time.Since(start)
	if !ok || elapsed < 50*time.Millisecond {
		t.Errorf("WaitForTimer should wait for the timer and return true")
	}

	count := CountTicks(10*time.Millisecond, 55*time.Millisecond)
	// Should get about 5-6 ticks depending on timing
	compare(t, true, count >= 4 && count <= 7, fmt.Sprintf("CountTicks should count ticks from the ticker, got %d", count))
}
