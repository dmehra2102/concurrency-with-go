package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

// This is problem of concurrency which is called RACE CONDITOION (when two or more operation run at th same time and try to update the same variable)

func main() {
	msg = "hello world"

	wg.Add(2)
	go updateMessage("Hello, Node!")
	go updateMessage("Hello, golang!")

	wg.Wait()
	fmt.Print(msg)
}
