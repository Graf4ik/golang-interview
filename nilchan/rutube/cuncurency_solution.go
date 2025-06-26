package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// программа должна запускать N горутин, которые выполняют функцию do
// так как программа должна подчитывать сколько было проведено секунд во сне, выводить результат каждую секунду
// так же определить, какая из горутин закончит выполнение первая

func do(id, dur int, done chan<- string) {
	sleepDuration := time.Duration(dur) * time.Second
	time.Sleep(sleepDuration)
	done <- fmt.Sprintf("Горутина #%d завершилась за %d секунд", id, dur)
}

func main() {
	const goroutineCount = 5

	// Канал для отслеживания завершения
	done := make(chan string, goroutineCount)

	// Генерируем случайные продолжительности
	durations := make([]int, goroutineCount)
	for i := 0; i < goroutineCount; i++ {
		durations[i] = rand.Intn(5) + 1
	}

	// Запускаем горутины
	for i := 0; i < goroutineCount; i++ {
		go do(i+1, durations[i], done)
	}

	// Таймер отслеживания
	ticker := time.NewTicker(1 * time.Second)
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		msg := <-done // только первое сообщение
		fmt.Println("🚀 Первая завершившаяся:", msg)
		wg.Done()
	}()

	go func() {
		for t := range ticker.C {
			elapsed := int(t.Sub(start).Seconds())
			fmt.Printf("⌛ Прошло %d секунд...\n", elapsed)
		}
	}()

	wg.Wait()
	ticker.Stop()
}
