package main

import (
	"log"
	"sync"
	"time"
)

var maxReq = 5

func main() {
	wg := sync.WaitGroup{}

	semaphore := make(chan struct{}, maxReq)

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(taskID int) {
			defer wg.Done()
			semaphore <- struct{}{} // резервируем место в семафоре перед началом работы

			time.Sleep(1 * time.Second)
			log.Printf("Запущен рабочий %d", taskID)

			<-semaphore // когда горутина завершает работу, освобождаем место и уменьшаем счетчик WaitGroup
		}(i)
	}

	wg.Wait()
}
