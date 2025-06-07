package livecoding_checklist

import (
	"fmt"
	"time"
)

// 22. Использование for (бесконечный цикл, проверка условия, проверка счетчика до значения вперед или назад)

func main() {
	// 1. Бесконечный цикл
	for {
		fmt.Println("Бесконечный цикл работает...")
		time.Sleep(1 * time.Second)

		// Выход по условию
		if time.Now().Second() == 0 {
			fmt.Println("Выход из цикла (секунда равна 0)")
			break
		}
	}
	count := 0

	// Цикл с условием (пока count < 5)
	for count < 5 {
		fmt.Printf("Итерация %d\n", count)
		count++
	}

	// Классический for с инициализацией, условием и инкрементом
	for i := 0; i < 5; i++ {
		fmt.Printf("Прямой счет: %d\n", i)
	}

	// Обратный счет
	for i := 5; i > 0; i-- {
		fmt.Printf("Обратный счет: %d\n", i)
	}

	// Цикл по слайсу
	fruits := []string{"яблоко", "банан", "апельсин"}
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}

	// Цикл по мапе
	ages := map[string]int{
		"Анна": 25,
		"Иван": 30,
		"Петр": 28,
	}
	for name, age := range ages {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// Пример с continue и break
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Пропускаем четные числа
		}
		if i > 7 {
			break // Выходим при i > 7
		}
		fmt.Println("Нечетное число:", i)
	}

	// Вложенные циклы с меткой для break
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				break outerLoop // Выход из обоих циклов
			}
		}
	}

	// Цикл с несколькими переменными
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}

	// Имитация do-while (выполнить хотя бы один раз)
	for {
		fmt.Println("Выполняем хотя бы один раз")
		count++

		if count >= 3 {
			break
		}
	}

	/*
		Практические примеры
		Обработка элементов с условием
	*/
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range numbers {
		if num%3 == 0 {
			fmt.Printf("Число %d (индекс %d) делится на 3\n", num, i)
		}
	}

	// Поиск элемента
	names := []string{"Alice", "Bob", "Charlie", "David"}
	found := false

	for _, name := range names {
		if name == "Charlie" {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Имя найдено")
	} else {
		fmt.Println("Имя не найдено")
	}
}

/*
Цикл for в Go - это универсальный инструмент, который может адаптироваться под различные сценарии использования.
Выбирайте подходящий вариант в зависимости от конкретной задачи.
*/
