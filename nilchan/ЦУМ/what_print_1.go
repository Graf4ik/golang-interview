package ЦУМ

import "fmt"

func main() {
	var m map[string]int

	for _, word := range []string{"hello", "world", "from", "the", "best", "language", "int", "the", "world"} {
		m[word]++
	}

	// Что будет выведено?
	for k, v := range m {
		fmt.Println(k, v) // panic: assignment to entry in nil map
	}

	fmt.Println(m) // Что будет выведено?
	// panic: assignment to entry in nil map
	// + есть ещё ошибка, нужно её найти
}
