package livecoding_checklist

import "fmt"

// 5. Итерация по слайсу (с индексом), мапе (итерация по ключам, значениям, использование ok).
// Доставание элементов по индексу или ключу. Добавление элементов.

func main() {
	// 1. Итерация по слайсу (с индексом)
	// Создаем слайс
	fruits := []string{"apple", "banana", "orange"}

	// Итерация с индексом
	for i, fruit := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, fruit)
	}

	// Только индексы
	for i := range fruits {
		fmt.Printf("Index: %d\n", i)
	}

	// Только значения (используем _ для индекса)
	for _, fruit := range fruits {
		fmt.Printf("Value: %s\n", fruit)
	}

	// 2. Доставание элементов слайса по индексу
	// Получение элемента
	first := fruits[0]  // "apple"
	second := fruits[1] // "banana"

	// Безопасное получение (проверка длины)
	if len(fruits) > 2 {
		third := fruits[2]
		fmt.Println(third)
	}

	// 3. Добавление элементов в слайс
	// Добавление одного элемента
	fruits = append(fruits, "grape")

	// Добавление нескольких элементов
	fruits = append(fruits, "pear", "peach")

	// Добавление другого слайса
	moreFruits := []string{"mango", "kiwi"}
	fruits = append(fruits, moreFruits...)

	// 4. Итерация по мапе
	// Создаем мапу
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 28,
	}

	// Итерация по ключам и значениям
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Только ключи
	for name := range ages {
		fmt.Println("Name:", name)
	}

	// Проверка существования ключа (использование ok)
	if age, ok := ages["Dave"]; ok {
		fmt.Println("Dave's age is", age)
	} else {
		fmt.Println("Dave not found")
	}

	// 5. Доставание элементов мапы по ключу
	// Получение значения
	aliceAge := ages["Alice"] // 25

	// Проверка существования ключа
	bobAge, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age:", bobAge)
	}

	// Идиоматичный способ
	if charlieAge, ok := ages["Charlie"]; ok {
		fmt.Println("Charlie's age:", charlieAge)
	} else {
		fmt.Println("Charlie not found")
	}

	// 6. Добавление и обновление элементов в мапе
	// Добавление нового элемента
	ages["Dave"] = 35

	// Обновление существующего элемента
	ages["Alice"] = 26

	// Добавление только если ключ не существует
	if _, exists := ages["Eve"]; !exists {
		ages["Eve"] = 28
	}

	// 7. Удаление элементов из мапы
	// Удаление элемента
	delete(ages, "Bob")

	// Безопасное удаление (не паникует, если ключа нет)
	delete(ages, "Unknown")

	// 8. Полный пример работы со слайсами и мапами
	// Работа со слайсом
	numbers := []int{10, 20, 30}

	// Итерация
	for i, num := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, num)
	}

	// Добавление
	numbers = append(numbers, 40, 50)
	fmt.Println("After append:", numbers)

	// Получение элемента
	fmt.Println("First element:", numbers[0])

	// Работа с мапой
	capitals := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
	}

	// Итерация
	for country, capital := range capitals {
		fmt.Printf("Capital of %s is %s\n", country, capital)
	}

	// Добавление
	capitals["Japan"] = "Tokyo"

	// Проверка и получение
	if capital, ok := capitals["Italy"]; ok {
		fmt.Println("Found Italy:", capital)
	}

	// Удаление
	delete(capitals, "France")
	fmt.Println("After delete:", capitals)
}

/*
9. Особенности работы
Слайсы:

Индексация начинается с 0
При обращении к несуществующему индексу - паника
append может изменить базовый массив (возвращает новый слайс)

Мапы:

Порядок итерации не гарантирован (случайный в Go)
При запросе несуществующего ключа возвращается нулевое значение типа
Используйте ok для проверки существования ключа
Мапы ссылочного типа - передаются по ссылке

Производительность:
Доступ по индексу в слайсе - O(1)
Доступ по ключу в мапе - O(1) в среднем случае
Итерация по всем элементам - O(n)
*/
