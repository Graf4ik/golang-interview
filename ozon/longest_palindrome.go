package main

import "fmt"

// для заданной строки необходимо определить длину самого большого плаиндрома, который
// можно составить из её символов.
// Пример: Input: aaabbbccccd Output: 11(палинлром dccbaaabccd)

func longestPalindromeLength(s string) int {
	charCount := make(map[rune]int)
	for _, ch := range s {
		charCount[ch]++
	}

	length := 0
	hasOdd := false
	for _, count := range charCount {
		length += count / 2 * 2 // добавляем чётное количество
		if count%2 == 1 {
			hasOdd = true // запомним, что можем поставить 1 символ в центр
		}
	}

	if hasOdd {
		length++ // центральный символ
	}
	return length
}

func main() {
	input := "aaabbbccccd"
	fmt.Println(longestPalindromeLength(input)) // Output: 11
}
