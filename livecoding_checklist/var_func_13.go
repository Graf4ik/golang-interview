package livecoding_checklist

import (
	"fmt"
	"sort"
)

// 13. Переменная с функцией. Вызов такой переменной с аргументами.

// 1. Простая функция без аргументов
func greet() {
	fmt.Println("Hello, World!")
}

// 2. Функция с аргументами и возвращаемым значением
func add(a, b int) int {
	return a + b
}

func main() {
	// 1. Простая функция без аргументов
	// Присваиваем функцию переменной
	myFunc := greet

	// Вызываем через переменную
	myFunc() // Выведет: Hello, World!

	// 2. Функция с аргументами и возвращаемым значением
	// Присваиваем функцию переменной
	mathOp := add

	// Вызываем с аргументами
	result := mathOp(3, 5)
	fmt.Println(result) // Выведет: 8

	// Для явного указания типа функции:
	// Объявляем тип функции
	type operation func(int, int) int

	// Присваиваем функцию переменной с явным типом
	var op operation = func(a, b int) int {
		return a * b
	}

	fmt.Println(op(4, 5)) // Выведет: 20

	// Анонимные функции
	// Можно сразу присваивать анонимную функцию:
	// Присваиваем анонимную функцию
	square := func(x int) int {
		return x * x
	}

	fmt.Println(square(5)) // Выведет: 25

	// Можно сразу вызвать
	cube := func(x int) int {
		return x * x * x
	}(3) // Вызываем сразу с аргументом 3

	fmt.Println(cube) // Выведет: 27

	// Функции как аргументы других функций
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	fmt.Println(applyOperation(2, 3, add))      // 5
	fmt.Println(applyOperation(2, 3, multiply)) // 6

	// Практический пример: кастомная сортировка
	people := []string{"Alice", "Bob", "Charlie"}

	// Присваиваем функцию сравнения
	byLength := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}

	// Используем в сортировке
	sort.Slice(people, byLength)

	fmt.Println(people) // [Bob Alice Charlie]

	// Возвращение функций из функций
	double := makeMultiplier2(2)
	triple := makeMultiplier2(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
}

func makeMultiplier2(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}
