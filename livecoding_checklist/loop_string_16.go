package livecoding_checklist

import "fmt"

// 16. Итерация по строке (через слайс и через ренж). Замена символа в середине строки (в том числе для строк с многобайтовыми символами).
func main() {
	// 1. Итерация по байтам (слайс рун)
	s := "Hello, 世界"

	// Итерация по байтам
	fmt.Println("By bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	// Вывод: 48 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c

	// 2. Итерация по рунам (range)
	s2 := "Hello, 世界"

	// Итерация по рунам (Unicode символам)
	fmt.Println("\nBy runes:")
	for i, r := range s2 {
		fmt.Printf("%d: %q (U+%04X)\n", i, r, r)
	}
	/*
	   Вывод:
	   0: 'H' (U+0048)
	   1: 'e' (U+0065)
	   2: 'l' (U+006C)
	   3: 'l' (U+006C)
	   4: 'o' (U+006F)
	   5: ',' (U+002C)
	   6: ' ' (U+0020)
	   7: '世' (U+4E16)
	   10: '界' (U+754C)
	*/

	// Замена символа в строке
	// 1. Для ASCII строк (однобайтовые символы)
	s3 := "Hello"
	s3 = replaceCharASCII(s3, 1, 'a')
	fmt.Println(s) // Hallo

	// 2. Для Unicode строк (многобайтовые символы)
	s4 := "Hello, 世界"
	s4 = replaceCharUnicode(s4, 7, '地') // Заменяем '世' на '地'
	fmt.Println(s)                      // Hello, 地界

	// 1. Преобразование строки в слайс рун
	s5 := "こんにちは" // Японское приветствие

	// Преобразуем в руны
	runes := []rune(s5)
	fmt.Println("Length in runes:", len(runes)) // 5
	fmt.Println("Length in bytes:", len(s5))    // 15

	// Модификация
	runes[2] = 'ニ' // Заменяем 3-й символ
	modified := string(runes)
	fmt.Println(modified) // こんニちは

	// 2. Нахождение позиции символа в байтах
	s6 := "Hello, 世界"
	pos := charBytePosition(s6, 7) // Позиция 7-го руны в байтах
	fmt.Println(pos)               // 7 (символ '世' начинается с 7-го байта)

	// Пример: замена всех вхождений символа
	s7 := "Hello, 世界 世界"
	s7 = replaceAllRunes(s7, '世', '天')
	fmt.Println(s7) // Hello, 天界 天界
}

/*
Важные замечания:

Строки в Go иммутабельны - любая модификация создает новую строку
len(s) возвращает длину в байтах, а не в символах
Для подсчета символов используйте utf8.RuneCountInString(s)
Для сложных операций со строками используйте пакет unicode/utf8
*/

/*
Работа со строками в Go, особенно с Unicode, требует понимания разницы между байтами и рунами.
Использование []rune для модификации строк - самый надежный способ для работы с многобайтовыми символами.
*/

func replaceAllRunes(s string, old, new rune) string {
	runes := []rune(s)
	for i, r := range runes {
		if r == old {
			runes[i] = new
		}
	}
	return string(runes)
}

func replaceCharASCII(s string, index int, newChar byte) string {
	if index < 0 || index >= len(s) {
		return s
	}
	bytes := []byte(s)
	bytes[index] = newChar
	return string(bytes)
}

func replaceCharUnicode(s string, index int, newChar rune) string {
	runes := []rune(s)
	if index < 0 || index >= len(runes) {
		return s
	}
	runes[index] = newChar
	return string(runes)
}

func charBytePosition(s string, runeIndex int) int {
	i := 0
	for n := range s {
		if i == runeIndex {
			return n
		}
		i++
	}
	return -1
}
