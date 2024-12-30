package main

import (
	"fmt"
	"sync"
)

var (
	count = 0
	mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer mu.Unlock()
	defer wg.Done()
	mu.Lock()
	count++

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", count)

	StructMutex()
}
