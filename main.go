package main

import (
	"fmt"
	"sync"
)

var Count int64 = 4

func main() {
	var wg sync.WaitGroup

	// WaitGroup example
	wg.Add(2)
	go WaitGroupFunc1(&wg)
	go WaitGroupFunc2(&wg)
	wg.Wait()

	fmt.Println("Hello World. We will gonna learn concurrency", Count)
}
