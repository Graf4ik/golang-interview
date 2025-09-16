package Ситидрайв

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1. что выведет
// 2. Mutex это? Какие тут примитивы синхронизации можем тут использовать?
// Что такое атомики? Напиши решение на атомиках. Напиши решение на каналах
// 3. Каналы это, чем буферезированный отличиается от небуферезированного
// 4. Что такое мапа, синк мапа? Чем отличается от мапы

func main() {
	data := map[int]bool{}
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		mu.Lock()
		data[i] = true
		mu.Unlock()
	}

	time.Sleep(1 * time.Second)

	for k, v := range data {
		fmt.Println(k, v) // Программа выведет 10 пар ключ-значение от 0 до 9, в произвольном порядке, например:
	}

	// 2
	var counter int32
	for i := 0; i < 10; i++ {
		atomic.AddInt32(&counter, 1)
	}
	fmt.Println("Final value:", atomic.LoadInt32(&counter))

	// 3
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
