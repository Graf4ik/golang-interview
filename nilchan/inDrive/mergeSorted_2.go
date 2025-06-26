package main

import (
	"fmt"
)

// Сложить два отсортированных массива

func main() {
	var a = []int{1, 2, 5}
	var b = []int{1, 2, 3, 4, 6}

	fmt.Println(mergeSorted(a, b))
}

func mergeSorted(first []int, second []int) []int {
	totalLen := len(first) + len(second)
	res := make([]int, 0, totalLen)

	r1 := 0
	r2 := 0

	for r1 < len(first) && r2 < len(second) {
		if first[r1] < second[r2] {
			res = append(res, first[r1])
			r1++
		} else {
			res = append(res, second[r2])
			r2++
		}
	}

	res = append(res, first[r1:]...)
	res = append(res, second[r2:]...)

	return res
}
