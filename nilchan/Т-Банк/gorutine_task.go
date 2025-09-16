package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Канал потокобезопасен (т.к. под капотом у него лежит sync.Mutex)
// Поэтому нет смысла объявлять еще один
func main() {
	var wg sync.WaitGroup
	m := make(chan string, 3)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(mm chan<- string, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			mm <- fmt.Sprintf("Gorutine %s", strconv.Itoa(i))
		}(m, i, &wg)
	}

	go func() {
		wg.Wait()
		close(m)
	}()

	for v := range m {
		fmt.Println(v)
	}

	// Deadlock из-за бесконечного цикла
	for {
		select {
		case q := <-m:
			fmt.Println(q)
		}
	}

	// Недостижимый код из-за бесконечного цикла
	// wg.Wait()
	// close(m)
}
