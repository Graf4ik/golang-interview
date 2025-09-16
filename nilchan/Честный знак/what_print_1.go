package main

import "fmt"

// Что выведет
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	foo(a[1:3])    // 2 3 len=2 cap=9
	println(a)     // println(a) — не выведет содержимое слайса, а выведет адрес (uintptr)
	fmt.Println(a) // [1 2 3 11 11 11 11 8 9 10]
}

func foo(a []int) {
	a = append(a, 11, 11, 11, 11)
}
