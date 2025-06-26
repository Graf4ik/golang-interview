package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5 // количество задач для выполнения
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("worker %d finished job %d\n", result, a)
	}
}

/*
Как работает этот код:
Очередь задач (jobs) содержит задачи для обработки. Она наполняется заданиями в основном потоке программы.
Пул воркеров — создаются три горутины, каждая из которых представляет рабочего. Они получают задачи из канала jobs, обрабатывают их и отправляют результаты в канал results.
Когда все задачи отправлены в очередь, канал задач закрывается с помощью close(jobs), что сигнализирует воркерам о завершении работы.
Результаты обрабатываются по мере поступления и выводятся в консоль.
*/

/*
Worker Pool отлично подходит для ситуаций, когда нужно обрабатывать большое количество однотипных задач,
например, входящие запросы к API, работа с файлами, запросы к базе данных и другие задачи, требующие параллельной обработки.
Паттерн позволяет эффективно распределять нагрузку, не создавая излишнего количества горутин.
*/
