package main

import "fmt"

// Что выведет
func main() {
	var numbers []*int

	for _, value := range []int{1, 2, 3, 4, 5} {
		value := value // версия до 1.19
		numbers = append(numbers, &value)
	}

	for _, number := range numbers {
		fmt.Printf("%d", *number)
	}
}
