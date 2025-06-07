package livecoding_checklist

import (
	"fmt"
	"time"
)

// 6. Создание каналов, с буфером и без. Закрытие канала. Запись и чтение из канала (range, select, напрямую, использование ok).
//Канал пустых структур, кастомных структур. Паттерн Fan in.

func main() {
	// 1. Создание каналов
	// Небуферизированный канал (размер 0)
	ch1 := make(chan int)

	// Буферизированный канал (размер 10)
	ch2 := make(chan string, 10)

	// Канал пустых структур (сигнальный)
	signalCh := make(chan struct{})

	// Канал кастомных структур
	type Message struct {
		Text string
		Code int
	}
	msgCh := make(chan Message)

	// 2. Запись и чтение из канала
	// Запись (в горутине, иначе deadlock для небуферизированного)
	go func() { ch1 <- 42 }()

	// Чтение
	value := <-ch1

	// Проверка на закрытие
	value, ok := <-ch1
	if !ok {
		fmt.Println("Channel closed")
	}

	// 3. Закрытие канала
	close(ch1)

	// Проверка при чтении
	val, ok := <-ch1
	if !ok {
		fmt.Println("Channel closed")
	}

	// 4. Итерация по каналу (range)
	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- fmt.Sprintf("Message %d", i)
		}
		close(ch2)
	}()

	for msg := range ch2 {
		fmt.Println(msg)
	}

	// 5. Использование select
	select {
	case v := <-ch1:
		fmt.Println("Received from ch1:", v)
	case ch2 <- "hello":
		fmt.Println("Sent to ch2")
	default:
		fmt.Println("No communication")
	}

	// 6. Канал пустых структур
	done := make(chan struct{})

	// Сигнализация завершения
	go func() {
		// Выполняем работу...
		close(done)
	}()

	// Ожидание завершения
	<-done

	// 8. Полный пример
	// Создаем каналы
	tasks := make(chan Task, 10)
	results := make(chan Task, 10)

	// Запускаем воркеров
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// Отправляем задачи
	for i := 1; i <= 5; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	// Получаем результаты
	for i := 1; i <= 5; i++ {
		res := <-results
		fmt.Printf("Received result: %+v\n", res)
	}

	// Пример Fan-in
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- fmt.Sprintf("ch1-%d", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- fmt.Sprintf("ch2-%d", i)
			time.Sleep(700 * time.Millisecond)
		}
	}()

	merged := fanIn(ch1, ch2)
	for i := 0; i < 6; i++ {
		fmt.Println("Fan-in received:", <-merged)
	}

	// Сигнальный канал
	done := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()
	<-done
	fmt.Println("Done signal received")
}

type Task struct {
	ID     int
	Result string
}

func worker(id int, tasks <-chan Task, results chan<- Task) {
	for t := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, t.ID)
		time.Sleep(time.Second) // Имитация работы
		t.Result = fmt.Sprintf("Result of task %d", t.ID)
		results <- t
	}
}

// 7. Паттерн Fan-in
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

/*
9. Важные особенности каналов

Небуферизированные каналы:
Отправка блокируется, пока другая горутина не готова принять
Прием блокируется, пока данные не будут отправлены

Буферизированные каналы:
Отправка блокируется только когда буфер полон
Прием блокируется только когда буфер пуст

Закрытие каналов:
Только отправитель должен закрывать канал
Отправка в закрытый канал вызывает панику
Чтение из закрытого канала возвращает нулевое значение и false

Select:
Позволяет ждать на нескольких каналах
Случай default выполняется, если другие не готовы

Fan-in:
Полезен для агрегации данных из нескольких источников
Может использоваться для реализации pub/sub систем

Каналы пустых структур:
Используются для сигнализации
Занимают минимальную память (0 байт)
*/
