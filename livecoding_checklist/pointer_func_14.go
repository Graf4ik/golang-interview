package livecoding_checklist

import "fmt"

// 14. Передача указателя в функцию. Модификация значения в указателе внутри функции.

/*
В Go все аргументы передаются по значению (копируются),
поэтому для модификации исходных переменных нужно использовать указатели.
*/

// 1. Передача указателя в функцию
func modifyValue2(ptr *int) {
	*ptr = 100 // Модифицируем значение по указателю
}

func main() {
	x := 42
	fmt.Println("Before:", x) // Before: 42

	modifyValue2(&x) // Передаем указатель на x

	fmt.Println("After:", x) // After: 100
}

// 2. Модификация структуры через указатель
type Person struct {
	Name string
	Age  int
}

func birthday(p *Person) {
	p.Age++ // Увеличиваем возраст
	// Эквивалентно: (*p).Age++
}

func main2() {
	bob := Person{"Bob", 25}
	fmt.Println("Before:", bob) // Before: {Bob 25}

	birthday(&bob)

	fmt.Println("After:", bob) // After: {Bob 26}
}

// Различные способы работы с указателями
// 1. Получение указателя на переменную

func main3() {
	a := 10
	ptr := &a // Получаем указатель

	fmt.Println("Value:", *ptr)  // Value: 10
	fmt.Println("Pointer:", ptr) // Pointer: 0xc000018030 (пример адреса)

	// 2. Создание указателя через new()
	ptr2 := new(int) // Создает int с нулевым значением и возвращает указатель
	*ptr2 = 42

	fmt.Println(*ptr2) // 42
}

// Практические примеры
// 1. Обмен значений через указатели
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main4() {
	x, y := 10, 20
	fmt.Println("Before:", x, y) // Before: 10 20

	swap(&x, &y)

	fmt.Println("After:", x, y) // After: 20 10
}

// 2. Модификация слайса (хотя слайсы уже содержат указатель)
func appendToSlice(s *[]int, values ...int) {
	*s = append(*s, values...)
}

func main5() {
	nums := []int{1, 2, 3}
	fmt.Println("Before:", nums) // Before: [1 2 3]

	appendToSlice(&nums, 4, 5, 6)

	fmt.Println("After:", nums) // After: [1 2 3 4 5 6]
}

/*
Особенности и рекомендации

Указатели и nil: Всегда проверяйте, что указатель не nil перед разыменованием
Методы с получателем-указателем:

func (p *Person) SetName(name string) {
    p.Name = name
}
Когда использовать указатели:
Для модификации больших структур (избегаем копирования)
Когда нужно изменить исходное значение
Для работы с интерфейсами и полиморфизмом

Когда не нужно использовать указатели:
Для маленьких структур (копирование дешевле)
Когда работаете с интерфейсом io.Reader и подобными
*/

// Пример с методами:

type Counter struct {
	value int
}

// Метод с получателем-указателем
func (c *Counter) Increment() {
	c.value++
}

func (c Counter) GetValue() int {
	return c.value
}

func main6() {
	c := Counter{}
	c.Increment()
	c.Increment()

	fmt.Println(c.GetValue()) // 2
}
