package МТС

import "fmt"

func main() {
	// Входные данные
	sales := []int{10, 20, 30, 40, 50}
	// Параметры обработки
	threshold := 30
	percent := 10

	// Канал входных данных
	input := make(chan int)

	// Запуск Pipeline
	go func() {
		defer close(input)
		for _, sale := range sales {
			input <- sale
		}
	}()

	// Pipeline: Фильтрация -> Увеличение -> Агрегация

	filtered := filter(input, threshold)
	increased := increase(filtered, percent)
	result := sum(increased)

	// Вывод результатов
	fmt.Printf("Итоговая сумма: %d\n", <-result)
}

// Шаг 1: Фильтрация данных
func filter(input chan int, threshold int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range input {
			if value >= threshold {
				out <- value
			}
		}
	}()

	return out
}

// Шаг 2: Увеличение значений на процент
func increase(input <-chan int, percent int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for value := range input {
			out <- value + value*percent/10
		}
	}()
	return out
}

// Шаг 3: Агрегация (подсчёт суммы)
func sum(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		total := 0
		for value := range input {
			total += value
		}
		out <- total
	}()
	return out
}
