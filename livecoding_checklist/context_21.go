package livecoding_checklist

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 21. Использование контекста с таймаутом и отменой. Мочь писать код который таймаутится и заканчивает исполнение через н секунд (через контекст с таймаутом и select)

/*
Контексты в Go предоставляют мощный механизм для управления временем выполнения операций и их отменой.
Вот как правильно использовать контексты с таймаутами и отменой.
*/

// Базовый пример с таймаутом

func longOperation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second): // Имитация долгой операции
		fmt.Println("Операция завершена успешно")
		return nil
	case <-ctx.Done(): // Срабатывает при отмене контекста
		fmt.Println("Операция отменена:", ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// Создаем контекст с таймаутом 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Важно вызывать cancel для освобождения ресурсов

	err := longOperation(ctx)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

// Пример с HTTP-запросом

func fetchWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func main13() {
	resp, err := fetchWithTimeout("https://example.com", 3*time.Second)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Статус код:", resp.StatusCode)
}

/* Продвинутое использование с несколькими горутинами */

func worker4(ctx context.Context, id int, ch chan<- string) {
	select {
	case <-time.After(time.Duration(id) * time.Second):
		ch <- fmt.Sprintf("Работник %d завершил работу", id)
	case <-ctx.Done():
		ch <- fmt.Sprintf("Работник %d отменен: %v", id, ctx.Err())
	}
}

func main14() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	results := make(chan string, 5)

	// Запускаем несколько работников
	for i := 1; i <= 5; i++ {
		go worker4(ctx, i, results)
	}

	// Собираем результаты
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}

/* Комбинирование таймаута и ручной отмены*/

func processData1(ctx context.Context, data []int) ([]int, error) {
	result := make([]int, 0, len(data))

	for _, item := range data {
		// Проверяем, не отменен ли контекст
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		// Имитация обработки
		time.Sleep(500 * time.Millisecond)
		result = append(result, item*2)
	}

	return result, nil
}

func main15() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Запускаем обработку в отдельной горутине
	done := make(chan struct{})
	var result []int
	var err error

	go func() {
		result, err = processData1(ctx, data)
		close(done)
	}()

	// Ждем завершения или таймаута
	select {
	case <-done:
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	case <-ctx.Done():
		fmt.Println("Превышено время ожидания:", ctx.Err())
	}
}

/*
Лучшие практики:

Всегда вызывайте cancel() функцию, чтобы освободить ресурсы
Проверяйте ctx.Err() в длительных операциях
Используйте http.NewRequestWithContext для HTTP-запросов
Для сложных сценариев комбинируйте context.WithTimeout и context.WithCancel
Передавайте контекст явно как первый аргумент функций
*/

/* Пример с дедлайном (абсолютное время) */

func main16() {
	// Устанавливаем дедлайн на 10 секунд вперед от текущего момента
	deadline := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Функция, которая проверяет контекст
	doWork := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Работа остановлена:", ctx.Err())
				return
			default:
				fmt.Println("Работа продолжается...")
				time.Sleep(1 * time.Second)
			}
		}
	}

	go doWork(ctx)

	// Ждем либо отмены, либо истечения времени
	<-ctx.Done()
	fmt.Println("Основная горутина: контекст завершен")
}
/*
	Контексты в Go - это мощный инструмент для управления временем жизни операций, особенно в распределенных системах.
	Они позволяют корректно останавливать вложенные операции и освобождать ресурсы при прерывании.
 */