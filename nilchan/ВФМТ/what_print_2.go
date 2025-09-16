package main

import "fmt"

// что выведет программа на 17 и 19 строчках
func main() {
	s := []int{0, 0}
	m := make(map[int]int, 2)
	changeSlice(s, 5)
	fmt.Println(s) // 5 0
	changeMap(m, 5)
	fmt.Println(m) // [0:5 1:5]
}

func changeSlice(s []int, v int) {
	s[0] = v         // 5 0
	s = append(s, v) // 5 0 5
	s[1] = v + 5     // 5 10 5
}

func changeMap(m map[int]int, v int) {
	m[0] = v
	m[1] = v
}

/*
map — ссылочный тип (reference-like), передаётся по ссылке.
Изменения внутри функции видны снаружи.
*/
