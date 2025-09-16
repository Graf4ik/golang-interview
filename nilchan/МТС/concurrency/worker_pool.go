package МТС

import (
	"fmt"
	"sync"
)

// Необходимо написать Worker Pool, чтобы выполнить параллельно numJobs
// заданий, используя numWorkers горутин. Эти горутины должны быть запущены
// один раз за всё время выполнения программы.
// Описание
// 1. Функция worker :
// На вход принимает:
// функцию f , которая выполняет задание,
// канал jobs для получения аргументов, канал results для записи результатов.
// Читает задания из канала jobs и записывает результаты выполнения f(job) в канал results .
// 2. Функция main :
// Запускает функцию worker в количестве numWorkers горутин.
// В качестве первого аргумента worker использует функцию multiplier .
// Записывает числа от 1 до numJobs в канал jobs .
// Решение
// Читает и выводит полученные значения из канала results, параллельно работе воркеров.

func worker(f func(int) int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- f(job)
	}
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}

	for i := 1; i < numWorkers; i++ {
		wg.Add(1)
		go worker(multiplier, jobs, results, &wg)
	}

	// Отправка заданий:
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result) // Так как f = x * 10, ожидаемый вывод: 10, 20, 30, 40, 50 (в произвольном порядке).
	}
}
