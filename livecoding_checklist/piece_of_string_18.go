package livecoding_checklist

import (
	"fmt"
	"unicode/utf8"
)

// 18. Доставание куска строки из середины по индексам (в том числе многобайтовых символов).
/*
Извлечение подстроки по индексам с учетом многобайтовых символов в Go

Для корректной работы с Unicode строками (содержащими многобайтовые символы)
нужно использовать специальные подходы, так как прямое обращение по индексам байтов может привести к некорректным результатам.
*/
func substring(s string, start, end int) string {
	runes := []rune(s)
	if start < 0 || end > len(runes) || start > end {
		return "" // или можно вернуть ошибку
	}
	return string(runes[start:end])
}

func substringEfficient(s string, start, end int) string {
	count := 0
	byteStart := -1
	byteEnd := -1

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if count == start {
			byteStart = i
		}
		if count == end {
			byteEnd = i
			break
		}
		i += size
		count++
	}

	if byteStart == -1 {
		return ""
	}
	if byteEnd == -1 {
		return s[byteStart:]
	}
	return s[byteStart:byteEnd]
}

func firstNRunes(s string, n int) string {
	runes := []rune(s)
	if n > len(runes) {
		n = len(runes)
	}
	return string(runes[:n])
}

func lastNRunes(s string, n int) string {
	runes := []rune(s)
	if n > len(runes) {
		n = len(runes)
	}
	return string(runes[len(runes)-n:])
}

func safeSubstring(s string, start, end int) string {
	runes := []rune(s)

	// Корректировка границ
	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	if start >= end {
		return ""
	}

	return string(runes[start:end])
}

func substringLarge(s string, start, end int) string {
	var result []rune
	count := 0

	for _, r := range s {
		if count >= start && count < end {
			result = append(result, r)
		}
		count++
		if count >= end {
			break
		}
	}

	return string(result)
}

func main() {
	s := "Hello, 世界!"

	// Извлекаем символы с 7 по 9 (включительно)
	sub := substring(s, 7, 9)
	fmt.Println(sub) // 世界

	// 2. С использованием utf8.DecodeRuneInString (более эффективный способ)
	s2 := "Привет, 世界!"
	sub2 := substringEfficient(s2, 1, 5)
	fmt.Println(sub2) // риве

	// Практические примеры
	//1. Извлечение первого N символов
	s3 := "こんにちは世界"                 // "Здравствуй, мир" по-японски
	fmt.Println(firstNRunes(s3, 5)) // こんにちは

	// 2. Извлечение последних N символов
	s4 := "Hello, 世界!"
	fmt.Println(lastNRunes(s4, 3)) // 界!

	// 3. Безопасное извлечение (с проверкой границ)
	s5 := "Пример строки с русскими буквами"
	fmt.Println(safeSubstring(s5, 3, 10))   // мер стр
	fmt.Println(safeSubstring(s5, -5, 100)) // вся строка
	fmt.Println(safeSubstring(s5, 10, 3))   // пустая строка

	// Оптимизация для больших строк
	//Для очень больших строк преобразование всей строки в слайс рун может быть неэффективным.
	//В этом случае лучше использовать итерационный подход:
	// Большая строка (для примера)
	largeStr := "Очень длинная строка с множеством символов..."
	fmt.Println(substringLarge(largeStr, 6, 15)) // длинная с
}

/*
Важные замечания
Индексация: В Go индексы в слайсах и строках начинаются с 0
Границы: Всегда проверяйте границы индексов, чтобы избежать паники
Производительность:
	[]rune(s) создает копию всех данных строки
	Итерационный подход может быть медленнее, но экономит память
Библиотечные функции: Для простых случаев можно использовать utf8.DecodeRuneInString

Для большинства случаев преобразование строки в слайс рун - это самый простой и читаемый способ работы с подстроками в Unicode строках.
Однако для обработки очень больших строк или в высоконагруженных приложениях стоит рассмотреть более оптимизированные подходы.
*/
