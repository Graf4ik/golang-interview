package main

import "fmt"

func appendLenWrong(numbers []*int) {
	size := len(numbers) // 3
	numbers = append(numbers, &size)
}

func appendLenGoodFirst(numbers *[]*int) {
	size := len(*numbers)
	*numbers = append(*numbers, &size)
}

func appendLenGoodSecond(numbers []*int) []*int {
	size := len(numbers)
	numbers = append(numbers, &size)
	return numbers
}

func main() {
	numbers := make([]*int, 0, 5)
	var number int
	for range 3 {
		number++
		numbers = append(numbers, &number)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
}
