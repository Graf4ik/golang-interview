package main

import "fmt"

// что выведет в консоль
func main() {
	originalSlice := make([]string, 0, 3)           // 0 3
	originalSlice = append(originalSlice, "A", "B") // 2 3
	foo(originalSlice)
	fmt.Println(originalSlice[:cap(originalSlice)]) // что выведет? // A B foo
	// Запись originalSlice[:cap(originalSlice)] означает создание нового среза,
	// который включает все элементы от 0 до ёмкости, даже если они ещё не были инициализированы явно.
}

func foo(input []string) []string {
	output := append(input, "foo") // A B foo // 3 3
	return output
}
