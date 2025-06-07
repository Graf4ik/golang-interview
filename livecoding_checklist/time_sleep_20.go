package livecoding_checklist

import (
	"fmt"
	"time"
)

// 20. Остановка исполнения через time.Sleep()

/*
Функция time.Sleep() из пакета time позволяет приостановить выполнение текущей горутины на указанный период времени.
*/

func main() {
	// Основное использование
	fmt.Println("Начало программы")

	// Приостановка на 2 секунды
	time.Sleep(2 * time.Second)

	fmt.Println("Прошло 2 секунды")

	// Различные варианты задержки
	// 1. Задержка в секундах
	fmt.Println("Ждем 3 секунды...")
	time.Sleep(3 * time.Second)
	fmt.Println("Готово!")

	// 2. Задержка в миллисекундах
	fmt.Println("Короткая пауза - 500 миллисекунд")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Пауза завершена")

	// 3. Задержка в минутах
	fmt.Println("Ожидаем 1 минуту...")
	time.Sleep(1 * time.Minute)
	fmt.Println("Минута истекла")

	// Практические примеры
	// 1. Таймер с обратным отсчетом
	countdown(5)

	// 2. Имитация длительной операции
	items := []string{"A", "B", "C", "D"}

	for _, item := range items {
		processItem(item)
	}

	// 3. Ограничение частоты запросов
	urls := []string{"api1.example.com", "api2.example.com", "api3.example.com"}

	for _, url := range urls {
		makeRequest(url)
		time.Sleep(1 * time.Second) // Ограничение: 1 запрос в секунду
	}

	// Пример с горутинами
	// Запускаем несколько горутин
	for i := 1; i <= 3; i++ {
		go worker3(i)
	}

	// Даем горутинам время на выполнение
	time.Sleep(3 * time.Second)
	fmt.Println("Все работники завершили задачи")

	// Прерывание сна
	// Если нужно прервать time.Sleep() досрочно, можно использовать каналы:
	interrupt := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(interrupt) // Прерываем через 1 секунду
	}()

	fmt.Println("Начало ожидания (2 секунды)")
	interrupted := sleepWithInterrupt(2*time.Second, interrupt)

	if interrupted {
		fmt.Println("Ожидание прервано!")
	} else {
		fmt.Println("Ожидание завершено полностью")
	}
}

func sleepWithInterrupt(duration time.Duration, interrupt <-chan struct{}) bool {
	select {
	case <-time.After(duration):
		return false // сработал таймер
	case <-interrupt:
		return true // получен сигнал прерывания
	}
}

func worker3(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Работник %d: задача %d\n", id, i)
		time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	}
}

func makeRequest(url string) {
	fmt.Printf("Запрос к %s\n", url)
	// Имитация HTTP-запроса
	time.Sleep(500 * time.Millisecond)
}

func processItem(item string) {
	fmt.Printf("Обработка %s...\n", item)
	time.Sleep(1 * time.Second) // Имитация работы
	fmt.Printf("%s обработан\n", item)
}

func countdown(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Время вышло!")
}

/*
Важные особенности

Горутины: time.Sleep() останавливает только текущую горутину, другие горутины продолжают работать
Точность: Гарантируется минимальная, но не максимальная длительность сна
Отрицательные значения: Если передать отрицательное время, функция немедленно вернет управление
Альтернативы: Для более сложных сценариев можно использовать time.Timer или time.Ticker
*/

/*
Функция time.Sleep() - это простой и эффективный способ добавить задержки в Go-программах,
но для сложных сценариев управления временем лучше использовать более гибкие механизмы из пакета time.
*/