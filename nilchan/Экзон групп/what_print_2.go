package main

import (
	"fmt"
	"sync"
)

var (
	a         = 0
	globalMap = map[string][]int{
		"test":  make([]int, 0),
		"test2": make([]int, 0),
		"test3": make([]int, 0),
	}
	mu = &sync.Mutex{}
)

// Какие проблемы есть в коде?
// Что такое race condition?
func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		wg.Done()
		mu.Lock()
		defer mu.Unlock()
		a = 10
		globalMap["test"] = append(globalMap["test"], a)
	}()

	go func() {
		wg.Done()
		mu.Lock()
		defer mu.Unlock()
		a = 11
		globalMap["test2"] = append(globalMap["test2"], a)
	}()

	go func() {
		wg.Done()
		mu.Lock()
		defer mu.Unlock()
		a = 12
		globalMap["test3"] = append(globalMap["test3"], a)
	}()

	wg.Wait()

	fmt.Printf("%v", globalMap)
	fmt.Printf("%d", a)
}
