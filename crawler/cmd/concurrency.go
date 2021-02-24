package main

import (
	"sync"
	"time"
)

// callConcurrent runs all functions in fns concurrently and returns the total elapsed time.
func callConcurrent(fns []func()) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()
	for i := range fns {
		fn := fns[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
	return time.Since(start)
}
