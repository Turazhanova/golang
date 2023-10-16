package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data map[string]string
	mu   sync.RWMutex
)

func init() {
	data = make(map[string]string)
}

func readData(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	return data[key]
}

func writeData(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	data[key] = value
}

func main() {
	go writeData("key1", "value1")
	go writeData("key2", "value2")

	go func() {
		for i := 0; i < 5; i++ {
			key := fmt.Sprintf("key%d", i)
			value := readData(key)
			fmt.Printf("Read: %s=%s\n", key, value)
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 1; i <= 2; i++ {
		key := fmt.Sprintf("key%d", i)
		value := readData(key)
		fmt.Printf("Read after write: %s=%s\n", key, value)
	}
}
