package lesson9

import (
	"sync"
	"testing"
	"time"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// TestConcurrentIncrementSum checks for race conditions and incorrect goroutine logic.
func TestConcurrentIncrementSum(t *testing.T) {
	// The correct sum of numbers from 1 to 100 is 5050.
	expected := int64(5050)
	// This will fail because of the loop variable capture bug and because wg.Wait() is missing.
	actual := ConcurrentIncrementSum(100)
	testutils.Compare(t, expected, actual, "ConcurrentIncrementSum has a race condition or goroutines are not being handled correctly.")
}

// TestChannelPipeline checks for deadlocks in a channel-based pipeline.
func TestChannelPipeline(t *testing.T) {
	numJobs := 5
	results := ChannelPipeline(numJobs)

	// We expect to receive 5 results. If the pipeline is deadlocked, we won't.
	for i := 0; i < numJobs; i++ {
		select {
		case _, ok := <-results:
			if !ok {
				t.Fatal("Pipeline channel was closed prematurely.")
				return
			}
			// In a fixed version, we would check the result value here.
		case <-time.After(2 * time.Second):
			t.Fatal("Test timed out. ChannelPipeline is likely deadlocked. Are you sending results and closing channels?")
		}
	}
}

// TestUpdateMapConcurrently checks for race conditions when incrementing a map value.
func TestUpdateMapConcurrently(t *testing.T) {
	m := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	increments := 1000
	wg.Add(2)

	go func() {
		defer wg.Done()
		UpdateMapConcurrently(m, "counter", increments, &mu)
	}()
	go func() {
		defer wg.Done()
		UpdateMapConcurrently(m, "counter", increments, &mu)
	}()

	wg.Wait()

	expected := 2 * increments
	actual := m["counter"]

	// This will fail because the race condition will result in lost increments.
	// The user must use the mutex in UpdateMapConcurrently to make this pass.
	testutils.Compare(t, expected, actual, "The final count is wrong due to a race condition. Use the mutex to protect the map access.")
}

// TestMultiplexChannels checks for deadlocks when reading from multiple channels.
func TestMultiplexChannels(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Send a value to ch2 first. A correct implementation should receive this.
	go func() { ch2 <- "from ch2" }()

	select {
	case result := <-func() chan string {
		// This wrapper is to call MultiplexChannels in a non-blocking way for the test.
		resChan := make(chan string, 1)
		go func() { resChan <- MultiplexChannels(ch1, ch2) }()
		return resChan
	}():
		testutils.Compare(t, "from ch2", result, "MultiplexChannels should have received the value from ch2.")
	case <-time.After(2 * time.Second):
		t.Fatal("Test timed out. MultiplexChannels is likely blocked waiting on the wrong channel. Use a select statement.")
	}
}
