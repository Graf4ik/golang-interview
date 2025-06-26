package main

import (
	"fmt"
	"sync"
)

// Этот код завершится зависанием (deadlock), потому что ты
// пытаешься читать из канала ch, который никогда не будет закрыт.
func main() {
	// Канал создаётся, но не закрывается в конце, а range ch ждёт до закрытия.
	ch := make(chan int) // 🔴 небуферизированный канал
	wg := &sync.WaitGroup{}
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}

	// Можно добавить отдельную горутину, которая ждёт завершения wg.Wait() и потом закрывает канал:
	// Закрываем канал после завершения всех горутин
	//go func() {
	//	wg.Wait()
	//	close(ch)
	//}()

	wg.Wait()
	var sum int
	for v := range ch { //  Проблема: ch не закрыт — range ch блокируется.
		sum += v
	}
	fmt.Println(sum)
}
