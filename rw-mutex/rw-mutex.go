package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Counter int64
	mu      sync.RWMutex
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.mu.Unlock()
	c.mu.Lock()
	c.Counter++
}

func (c *Counter) Value() int64 {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.Counter
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Write operation
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	// Read Operation
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Current counter value:", counter.Value())
			time.Sleep(50 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}
