package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Counter int64
	mu      sync.Mutex
}

func (c *Counter) Value() int64 {
	return c.Counter
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer c.mu.Unlock()
	defer wg.Done()
	c.mu.Lock()
	c.Counter++
}

func StructMutex() {
	var wg sync.WaitGroup

	counter := Counter{
		Counter: 0,
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}
