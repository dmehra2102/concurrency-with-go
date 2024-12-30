package main

import (
	"sync"
)

func WaitGroupFunc1(wg *sync.WaitGroup) {
	// defer wg.Done()

	for i := 0; i < 3; i++ {
		Count++
	}
}

func WaitGroupFunc2(wg *sync.WaitGroup) {
	// defer wg.Done()

	for i := 0; i < 3; i++ {
		Count--
	}
}
