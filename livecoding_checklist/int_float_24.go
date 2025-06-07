package livecoding_checklist

import (
	"fmt"
	"strconv"
)

// 24. Конвертация строки в int или float.

/* В Go для преобразования строк в числовые типы используется пакет strconv. Вот основные методы конвертации. */

func main() {
	// 1. Конвертация строки в целое число (int)
	s := "42"

	// Конвертация в int (базовый 10)
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}
	fmt.Printf("Тип: %T, Значение: %d\n", num, num)
	// Вывод: Тип: int, Значение: 42

	/* С указанием системы счисления */
	s2 := "1010" // двоичное число

	// Конвертация из строки в int с указанием системы счисления (2) и размера (64 бита)
	num2, err2 := strconv.ParseInt(s2, 2, 64)
	if err2 != nil {
		fmt.Println("Ошибка конвертации:", err2)
		return
	}
	fmt.Printf("Десятичное значение: %d\n", num2)
	// Вывод: Десятичное значение: 10

	// 2. Конвертация строки в число с плавающей точкой (float)
	s3 := "3.1415"

	// Конвертация в float64
	num3, err := strconv.ParseFloat(s3, 64)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}
	fmt.Printf("Тип: %T, Значение: %f\n", num3, num3)
	// Вывод: Тип: float64, Значение: 3.141500

	// Специальные значения
	values := []string{"1.23", "-12.34", "NaN", "Inf", "+Inf"}

	for _, s := range values {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("%s: ошибка (%v)\n", s, err)
		} else {
			fmt.Printf("%s → %f\n", s, num)
		}
	}
	/*
		Вывод:
		1.23 → 1.230000
		-12.34 → -12.340000
		NaN → NaN
		Inf → +Inf
		+Inf → +Inf
	*/

	// 3. Обратная конвертация (числа в строку)
	num4 := 42

	// Конвертация int в строку
	s4 := strconv.Itoa(num4)
	fmt.Printf("Тип: %T, Значение: %q\n", s4, s4)
	// Вывод: Тип: string, Значение: "42"

	// Float в строку
	f := 3.1415

	// Конвертация float в строку с контролем формата
	s5 := strconv.FormatFloat(f, 'f', 2, 64) // 'f' - формат, 2 - знаков после точки, 64 - битность
	fmt.Printf("Тип: %T, Значение: %q\n", s5, s5)
	// Вывод: Тип: string, Значение: "3.14"

	/* 5. Альтернативные методы (fmt.Sscanf) */
	s6 := "42"
	var num int

	// Использование fmt.Sscanf для парсинга
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Число:", num) // Число: 42
}

/* 4. Продвинутые примеры */
func getUserInput() (int, error) {
	var input string
	fmt.Print("Введите число: ")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("некорректное число: %v", err)
	}
	return num, nil
}

func main21() {
	num, err := getUserInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Вы ввели:", num)
}

// Конвертация с проверкой диапазона
func safeAtoi(s string) (int, error) {
	// Сначала пробуем ParseInt для проверки переполнения
	num, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func main22() {
	bigNum := "2147483648" // > math.MaxInt32

	num, err := safeAtoi(bigNum)
	if err != nil {
		fmt.Println("Ошибка:", err) // Выведет ошибку переполнения
	}
}

/*
Рекомендации:

Всегда проверяйте ошибки при конвертации
Для целых чисел используйте strconv.Atoi для десятичных чисел
Для чисел в других системах счисления - strconv.ParseInt
Для чисел с плавающей точкой - strconv.ParseFloat
Для максимальной производительности в критических участках кода рассмотрите использование специализированных парсеров

Пакет strconv предоставляет все необходимые инструменты для надежной конвертации между строками и числовыми типами в Go.
*/
