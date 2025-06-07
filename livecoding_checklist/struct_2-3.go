package livecoding_checklist

import "fmt"

// 2. Объявлять структуру с полями разных типов, уметь ее инстанциировать в коде.
func main() {
	// Объявляем структуру Person
	type Person struct {
		Name     string
		Age      int
		Height   float64
		IsActive bool
		Scores   []int
		Metadata map[string]string
		Address  struct {
			City    string
			Country string
		}
	}

	var p1 Person // Все поля получат нулевые значения

	// Вариант 2: Литеральная инициализация
	p2 := Person{
		Name:   "Alice",
		Age:    30,
		Height: 1.75,
	}

	// Вариант 3: Позиционная инициализация (порядок важен)
	p3 := Person{"Bob", 25, 1.80, true, nil, nil, struct{ City, Country string }{}}

	// Вариант 4: Создание указателя на структуру
	p4 := &Person{
		Name: "Charlie",
		Age:  35,
	}

	// Вариант 5: С new()
	p5 := new(Person) // Возвращает указатель (*Person)

	// 3. Пример с вложенными структурами
	type Address struct {
		Street  string
		City    string
		ZipCode string
	}

	type Employee struct {
		ID        int
		Name      string
		Position  string
		Salary    float64
		Address   Address
		Projects  []string
		IsManager bool
	}

	// Инициализация
	emp := Employee{
		ID:       1001,
		Name:     "John Doe",
		Position: "Developer",
		Salary:   75000.50,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "10001",
		},
		Projects:  []string{"Project A", "Project B"},
		IsManager: false,
	}

	// 4. Анонимные структуры
	temp := struct {
		Key   string
		Value int
	}{
		Key:   "priority",
		Value: 10,
	}

	// 5. Доступ к полям структуры
	// Чтение
	name := p2.Name
	city := emp.Address.City

	// Запись
	p2.Age = 31
	emp.Position = "Senior Developer"

	/*
		6. Особенности структур в Go

		Экспортируемость: Поля с заглавной буквы экспортируются (видны в других пакетах)
		Теги структур: Можно добавлять метаданные к полям
		Сравнение: Структуры можно сравнивать, если все их поля сравниваемы
		Встраивание: Поддержка композиции через встраивание структур
	*/
	type User struct {
		Name string `json:"name" db:"user_name"`
		Age  int    `json:"age" validate:"min=18"`
	}

	type Book struct {
		Title       string
		Author      string
		Pages       int
		Chapters    []string
		Ratings     map[string]int
		IsHardcover bool
	}

	// Создание экземпляра книги
	book1 := Book{
		Title:       "The Go Programming Language",
		Author:      "Alan A. A. Donovan & Brian W. Kernighan",
		Pages:       380,
		IsHardcover: true,
		Chapters:    []string{"Tutorial", "Program Structure", "Basic Data Types"},
		Ratings: map[string]int{
			"Goodreads": 4,
			"Amazon":    5,
		},
	}

	// Создание с нулевыми значениями
	var book2 Book
	book2.Title = "Effective Go"
	book2.Author = "Google Go Team"

	// Вывод информации
	fmt.Printf("Book 1: %+v\n", book1)
	fmt.Printf("Book 2: %+v\n", book2)

	// Доступ к полям
	fmt.Println("First chapter of book 1:", book1.Chapters[0])
}
