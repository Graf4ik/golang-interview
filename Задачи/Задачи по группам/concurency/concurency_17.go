package main

import (
	"fmt"
	"sync"
	"time"
)

//===========================================================
//Задача 17
//1. Что выведется? Исправь проблему
//===========================================================

// # Вариант1
// ----------
func main() {
	x := make(map[int]int, 1)
	mu := sync.Mutex{}

	go func() {
		mu.Lock()
		defer mu.Unlock()
		x[1] = 2
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		x[1] = 7
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		x[1] = 10
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("x[1] =", x[1])
}
