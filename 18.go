package main

import (
	"fmt"
	"sync"
)

// Counter struct to hold count and mutex
type Counter struct {
	mu    sync.Mutex // Mutex to synchronize access to count
	count int        // Count of increments
}

// Increment function to increment the count
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensure mutex is unlocked when function exits
	c.count++
}

// GetCount function to return the current count
func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensure mutex is unlocked when function exits
	return c.count
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final count:", counter.GetCount())
}
