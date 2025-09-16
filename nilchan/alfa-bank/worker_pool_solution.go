package alfaBank

import (
	"fmt"
	"sync"
	"time"
)

// реализовать worker pool
// Есть 10 задач (функций), каждая засыпает на 1 сек и выводит номер воркера, который эту задачу исполнил.
// Кол-во воркеров задается при запуске.

func main() {
	const (
		numJobs    = 10
		numWorkers = 3
	)

	var wg sync.WaitGroup
	jobs := make(chan int)

	for w := 0; w < numWorkers; w++ {
		go func(workerId int) {
			for job := range jobs {
				fmt.Printf("worker %d started job %d\n", workerId, job)
				time.Sleep(1 * time.Second)
				fmt.Printf("worker %d finished job %d\n", workerId, job)
				wg.Done()
			}
		}(w)
	}

	for j := 0; j < numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}
	wg.Wait()
	close(jobs)
}
