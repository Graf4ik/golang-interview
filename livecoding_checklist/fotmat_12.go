package livecoding_checklist

import "fmt"

// 12. Форматирование строки. Вывод инта, флоата, була, кастомной структуры в форматированной строке.
/*
Основные функции форматирования

Printf - форматированный вывод
Sprintf - возвращает форматированную строку
Fprintf - запись форматированной строки в io.Writer
*/

func main() {
	/* Целые числа (int) */
	num := 42

	// Десятичное представление
	fmt.Printf("Decimal: %d\n", num) // Decimal: 42

	// Шестнадцатеричное
	fmt.Printf("Hex: %x\n", num)          // Hex: 2a
	fmt.Printf("Hex with 0x: %#x\n", num) // Hex with 0x: 0x2a

	// Двоичное
	fmt.Printf("Binary: %b\n", num) // Binary: 101010

	// С ведущими нулями
	fmt.Printf("Padded: %04d\n", num) // Padded: 0042

	/* Числа с плавающей точкой (float) */
	pi := 3.1415926535

	// По умолчанию 6 знаков после точки
	fmt.Printf("Default: %f\n", pi) // Default: 3.141593

	// С указанием точности
	fmt.Printf("2 digits: %.2f\n", pi) // 2 digits: 3.14

	// Научная нотация
	fmt.Printf("Scientific: %e\n", pi) // Scientific: 3.141593e+00

	// Большая ширина поля
	fmt.Printf("Width 10: %10.3f\n", pi) // Width 10:      3.142

	/* Логические значения (bool) */
	flag := true

	// Простое представление
	fmt.Printf("Boolean: %t\n", flag) // Boolean: true

	// С форматированием
	fmt.Printf("Boolean: %v\n", flag)  // Boolean: true
	fmt.Printf("Boolean: %#v\n", flag) // Boolean: true

	/* Форматирование структур и пользовательских типов
	Простая структура */
	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Alice", 30}

	// По умолчанию
	fmt.Printf("%v\n", p) // {Alice 30}

	// С именами полей
	fmt.Printf("%+v\n", p) // {Name:Alice Age:30}

	// Go-синтаксис
	fmt.Printf("%#v\n", p) // main.Person{Name:"Alice", Age:30}

	/* Пользовательский формат с интерфейсом Stringer */
	type Point struct {
		X, Y int
	}

	pt := Point{10, 20}
	fmt.Printf("Point: %v\n", pt) // Point: (10, 20)
	fmt.Printf("Point: %s\n", pt) // Point: (10, 20)

	/* Комплексное форматирование */
	name := "Bob"
	age := 25
	score := 87.5
	passed := true

	// Комбинированный вывод
	fmt.Printf(
		"Name: %-10s | Age: %03d | Score: %5.2f | Passed: %t\n",
		name, age, score, passed,
	)
	// Name: Bob        | Age: 025 | Score: 87.50 | Passed: true
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

/*
Полезные спецификаторы форматирования
%v - значение в формате по умолчанию
%#v - Go-синтаксическое представление значения
%T - тип значения
%% - знак процента (литерал)
*/
