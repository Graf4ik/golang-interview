package main

import "fmt"

func main() {
	s := "test"
	s1 := []rune(s)
	fmt.Println(s1[0]) // что выведет программа

	/*
		→ s1[0] — это руна 't'
		→ числовое значение: 116 (ASCII код)
	*/

	s1[0] = rune("R")
	fmt.Println(s1) // что выведет программа cannot convert "R" (untyped string constant) to type rune
}
