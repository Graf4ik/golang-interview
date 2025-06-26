package cloud

import (
	"fmt"
	"sync"
	"time"
)

/*
Есть некий сервис, который принимает запросы (TCP, HTTP - не имеет значения)
Клинет передаёт на вход некоторый объект с описанием задачи

Мы должны, получив эту задачу, поставить ее в очередь на обработку
Задача запускается в обработку (имитируем полезную работу через time.Sleep(5*time.Second)),
если у нас есть свободные обработчики.
Как только очередная задача выполнилась - берём следующую задачу из очереди. Если в очереди
пусто, ожидаем новых задач от клиентов.

Сервис одновременно может обрабатывать не более N задач. Остальные задачи должны помещаться в очередь.
*/

type Task struct {
	ID int
}

const (
	MaxConcurrentTasks = 3 // Максимум одновременных задач
)

func main() {
	taskQueue := make(chan Task, 100)
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, MaxConcurrentTasks)

	go func() {
		for task := range taskQueue {
			semaphore <- struct{}{}
			wg.Add(1)

			go func(t Task) {
				defer func() {
					<-semaphore
					wg.Done()

				}()
				fmt.Printf("Начинаем обработку задачи #%d\n", t.ID)
				time.Sleep(5 * time.Second) // имитация полезной работы
				fmt.Printf("Завершили задачу #%d\n", t.ID)
			}(task)
		}
	}()

	// Имитация поступающих от клиентов задач
	go simulateClientInput(taskQueue)
	close(taskQueue)
	wg.Wait()
	fmt.Println("Все задачи обработаны")
}

func simulateClientInput(queue chan<- Task) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Клиент отправил задачу #%d\n", i)
		queue <- Task{ID: i}
		time.Sleep(1 * time.Second)
	}
}
