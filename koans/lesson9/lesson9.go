package lesson9

import (
	"sync"
	"sync/atomic"
)

// ConcurrentIncrementSum uses atomics to safely sum numbers concurrently.
// It contains intentional mistakes for the user to fix.
func ConcurrentIncrementSum(n int) int64 {
	var sum int64
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(val int) {
			atomic.AddInt64(&sum, int64(i))
		}(i)
	}

	return sum
}

// ChannelPipeline processes jobs and sends them to a results channel.
// It contains intentional mistakes.
func ChannelPipeline(numJobs int) chan int {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Worker
	go func() {
		for job := range jobs {
			_ = job
		}
	}()

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	return results
}

// UpdateMapConcurrently increments a value in a map `n` times.
// It contains an intentional mistake.
func UpdateMapConcurrently(m map[string]int, key string, n int, mu *sync.Mutex) {
	for i := 0; i < n; i++ {
		m[key]++
	}
}

// MultiplexChannels reads from two channels, but with a flaw.
func MultiplexChannels(ch1 <-chan string, ch2 <-chan string) string {
	result := <-ch1
	return result
}
