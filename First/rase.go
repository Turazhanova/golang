package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func increment() {
	for i := 0; i < 1000; i++ {
		counter++
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}
