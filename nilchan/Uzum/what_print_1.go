package Uzum

import "fmt"

// что выведет
func main() {
	s := "hello"
	s[0] = 'H'
	fmt.Println(s)
}

// Этот код не скомпилируется, потому что строки (string) в Go иммутабельны (нельзя изменить символ по индексу).
// В Go строка — это неизменяемая последовательность байтов. Попытка изменить s[0] нарушает это правило.
