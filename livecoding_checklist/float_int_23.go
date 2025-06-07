package livecoding_checklist

import (
	"fmt"
	"reflect"
)

// 23. Каст обьектов друг к другу. Float -> Int и обратно. Каст обьектов одного интерфейса к конкретным типам. Разные виды кастов.

/*В Go существует несколько способов приведения типов, которые можно разделить на несколько категорий. */

// Float ↔ Int
func main() {
	// 1. Приведение базовых числовых типов
	// Float → Int (отбрасывается дробная часть)
	f := 3.14
	i := int(f)    // Явное приведение
	fmt.Println(i) // 3

	// Int → Float
	j := 42
	g := float64(j) // Явное приведение
	fmt.Println(g)  // 42.0

	// int32 → int64
	var a int32 = 1000
	b := int64(a)
	fmt.Printf("%T: %v\n", b, b) // int64: 1000

	// int → uint (может привести к неожиданным результатам для отрицательных чисел)
	c := -1
	d := uint(c)
	fmt.Printf("%T: %v\n", d, d) // uint: 18446744073709551615 (на 64-битной системе)

	// 2. Приведение интерфейсов к конкретным типам
	// Type Assertion (утверждение типа)
	var val interface{} = "hello"

	// Безопасное приведение
	if s, ok := val.(string); ok {
		fmt.Println("Это строка:", s) // Это строка: hello
	} else {
		fmt.Println("Это не строка")
	}

	// Небезопасное приведение (вызовет панику, если тип не совпадает)
	// s := val.(string)
	// fmt.Println(s)

	// 4. Приведение указателей
	type MyInt int
	var x int = 42
	var px *int = &x

	// Приведение указателей
	var py *MyInt = (*MyInt)(px)
	fmt.Println(*py) // 42

	// 5. Приведение с помощью рефлексии (reflect)
	var x2 float64 = 3.4

	// Получаем reflect.Value
	v := reflect.ValueOf(x2)

	// Пытаемся привести к int
	if v.CanInt() {
		i := v.Int()
		fmt.Println(i)
	} else {
		fmt.Println("Нельзя привести к int")
	}

	// Приведение через интерфейс
	y := v.Interface().(float64)
	fmt.Println(y) // 3.4
}

// Type Switch (переключатель типов)
func checkType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("Это int: %v\n", v)
	case float64:
		fmt.Printf("Это float64: %v\n", v)
	case string:
		fmt.Printf("Это string: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

func main17() {
	checkType(42)      // Это int: 42
	checkType(3.14)    // Это float64: 3.14
	checkType("hello") // Это string: hello
	checkType(true)    // Неизвестный тип: bool
}

// 3. Приведение между пользовательскими типами
// Базовые типы
type Celsius float64
type Fahrenheit float64

func main18() {
	var c Celsius = 100
	f := Fahrenheit(c*9/5 + 32)           // Явное приведение между пользовательскими типами
	fmt.Printf("%.2f°C = %.2f°F\n", c, f) // 100.00°C = 212.00°F
}

// Приведение структур
type Person4 struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Job  string
}

func main19() {
	p := Person4{"Alice", 30}

	// Нельзя напрямую привести Person к Employee
	// e := Employee(p) // Ошибка компиляции

	// Нужно явно преобразовывать
	e := Employee{
		Name: p.Name,
		Age:  p.Age,
		Job:  "Developer",
	}
	fmt.Println(e) // {Alice 30 Developer}
}

/*
Важные замечания:

В Go нет автоматического приведения типов (implicit conversion)
Все приведения должны быть явными
При работе с интерфейсами используйте type assertion или type switch
Будьте осторожны с приведением числовых типов - возможна потеря точности или неожиданные результаты
Для сложных преобразований между структурами часто лучше использовать специальные функции-конвертеры
*/

// Практический пример: безопасное приведение интерфейса
func safeStringConvert(val interface{}) (string, error) {
	switch v := val.(type) {
	case string:
		return v, nil
	case fmt.Stringer:
		return v.String(), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%v", v), nil
	default:
		return "", fmt.Errorf("неподдерживаемый тип: %T", val)
	}
}

func main20() {
	res, err := safeStringConvert(42)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(res) // 42
	}
}
