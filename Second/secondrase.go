package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var data []int

func appendData(value int) {
	mu.Lock()
	data = append(data, value)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			appendData(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Data:", data)
}
