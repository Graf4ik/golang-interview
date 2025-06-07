package livecoding_checklist

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 7. Создание горутины, возврат данных из горутины в главный поток. Использование mutex, wait group.

func main() {
	// 1. Создание горутины
	go func() {
		fmt.Println("Эта функция выполняется в отдельной горутине")
	}()

	// 2. Возврат данных из горутин c использованием каналов:
	resultCh := make(chan int)

	go func() {
		// Долгие вычисления
		result := 42
		resultCh <- result
	}()

	// Блокирующее чтение из канала
	result := <-resultCh
	fmt.Println("Получен результат:", result)

	// 3. Использование WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Увеличиваем счетчик
		go worker(i, wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Все горутины завершили работу")
	// 4. Использование Mutex
	counter := SafeCounter{}

	for i := 0; i < 1000; i++ {
		go counter.Increment()
	}
	time.Sleep(time.Second)
	fmt.Println("Итоговое значение:", counter.value)

	// 5. Комбинированный пример с возвратом данных
	const numWorkers = 5
	var wg sync.WaitGroup
	resultCh := make(chan int, numWorkers)

	// Запускаем горутины
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go processData(i, resultCh, &wg)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Собираем результаты
	for result := range resultCh {
		fmt.Println("Получен результат:", result)
	}
	// 6. Паттерн worker pool
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Создаем пул воркеров
	for w := 1; w <= numWorkers; w++ {
		go workerPool(w, jobs, results)
	}

	// Отправляем задания
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Получаем результаты
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Результат:", <-results)
	}

	// 8. Пример с контекстом для отмены
	ctx, cancel := context.WithCancel(context.Background())
	resultCh := make(chan int)

	wg.Add(1)
	go workerContext(ctx, &wg, resultCh)

	// Имитируем отмену через 1 секунду
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	select {
	case res := <-resultCh:
		fmt.Println("Результат:", res)
	case <-ctx.Done():
		fmt.Println("Операция прервана")
	}

	wg.Wait()
	close(resultCh)
}

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func processData(data int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Имитация обработки
	time.Sleep(time.Duration(data) * time.Millisecond)
	result := data * 2
	resultCh <- result
}

func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик при завершении
	fmt.Printf("Воркер %d начал работу\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d завершил работу\n", id)
}

func workerPool(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Воркер %d начал задачу %d\n", id, job)
		time.Sleep(time.Second) // Имитация работы
		results <- job * 2
		fmt.Printf("Воркер %d завершил задачу %d\n", id, job)
	}
}

func workerContext(ctx context.Context, wg *sync.WaitGroup, resultCh chan<- int) {
	defer wg.Done()

	select {
	case <-time.After(2 * time.Second):
		resultCh <- 42
	case <-ctx.Done():
		fmt.Println("Работа отменена")
		return
	}
}

/*
7. Важные особенности:

Горутины:
Легковесные потоки выполнения
Запускаются с ключевым словом go
Нет гарантии порядка выполнения

WaitGroup:
Add() - увеличивает счетчик
Done() - уменьшает счетчик
Wait() - блокирует выполнение, пока счетчик не станет 0

Mutex:
Lock()/Unlock() - защищают критическую секцию
Всегда используйте defer для разблокировки
RWMutex для разделения блокировок чтения/записи

Каналы:
Основной способ коммуникации между горутинами
Всегда закрывайте каналы на стороне отправителя
Используйте буферизированные каналы для уменьшения блокировок

Рекомендации:
Всегда проверяйте, закрыт ли канал (используя ok)
Используйте context для отмены операций
Избегайте утечек горутин (всегда предусматривайте способ завершения)
*/
