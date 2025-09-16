package Магнит

import "fmt"

// Что выведется?
func main() {
	x := 10
	y := 20

	defer func(val int) {
		fmt.Printf("x", val)
	}(x)

	defer func() {
		fmt.Printf("y", y)
	}()

	x = 100
	y = 200
} // y%!(EXTRA int=200)x%!(EXTRA int=10)

// Произошёл из-за отсутствия плейсхолдеров в форматных строках Printf().
// Go не ругается ошибкой, но пишет в stdout "диагностическое сообщение" о неправильном использовании.
