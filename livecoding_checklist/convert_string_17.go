package livecoding_checklist

import (
	"fmt"
	"strings"
)

// 17. Конвертация строку в слайс рун и обратно.
func main() {
	s := "Привет, 世界!" // Строка с кириллицей и иероглифами

	// Конвертация строки в слайс рун
	runes := []rune(s)

	fmt.Println("Слайс рун:", runes)
	// Вывод: [1055 1088 1080 1074 1077 1090 44 32 19990 30028 33]

	fmt.Println("Количество символов (рун):", len(runes))
	// Вывод: 11

	// Способ 2: Итерация с помощью range
	runes2 := stringToRunes(s)
	fmt.Println(runes2) // [72 101 108 108 111 44 32 19990 30028]

	// Преобразование слайса рун обратно в строку
	// Простое преобразование
	runes3 := []rune{'П', 'р', 'и', 'в', 'е', 'т', ' ', '世', '界'}

	// Конвертация слайса рун в строку
	s2 := string(runes3)

	fmt.Println("Результат:", s2)
	// Вывод: Привет 世界

	// Эффективное преобразование (с использованием strings.Builder)
	runes4 := []rune{'G', 'o', ' ', '語', '言'}
	s3 := runesToString(runes4)
	fmt.Println(s3) // Go 語言

	// Практические примеры
	// 1. Реверс строки с Unicode символами
	s4 := "Hello, 世界"
	reversed := reverseString(s4)
	fmt.Println(reversed) // 界世 ,olleH

	// 2. Подсчет количества определенных символов
	s5 := "банан"
	count := countRunes(s5, 'а')
	fmt.Println(count) // 2
}

/*
Особенности и рекомендации
Производительность: Преобразование в []rune создает новое выделение памяти, поэтому для больших строк учитывайте накладные расходы.

Длина строки:

len(s) - возвращает длину в байтах
len([]rune(s)) - возвращает количество Unicode символов

Модификация строк:
s := "hello"
runes := []rune(s)
runes[0] = 'H'
s = string(runes) // "Hello"

Оптимизация: Если вам нужно только подсчитать символы, используйте utf8.RuneCountInString(s) вместо полного преобразования в []rune.
Опасность: Не используйте индексацию строки как массива (s[0]), так как это даст вам байты, а не символы.

Работа со слайсами рун - это основной способ корректной обработки Unicode строк в Go,
особенно когда требуется модификация отдельных символов или работа с позициями символов в строке.
*/

func countRunes(s string, target rune) int {
	runes := []rune(s)
	count := 0
	for _, r := range runes {
		if r == target {
			count++
		}
	}
	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func runesToString(runes []rune) string {
	var sb strings.Builder
	sb.Grow(len(runes)) // Оптимизация: заранее выделяем память
	for _, r := range runes {
		sb.WriteRune(r)
	}
	return sb.String()
}

func stringToRunes(s string) []rune {
	runes := make([]rune, 0, len(s)) // Предварительное выделение памяти
	for _, r := range s {
		runes = append(runes, r)
	}
	return runes
}
