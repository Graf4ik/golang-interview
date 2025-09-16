package main

import "fmt"

// что и в каком порядке выведет программа на 11 и 14 строчках
func main() {
	var (
		i int = 1
		j int = 2
	)
	defer fmt.Println(i)
	i = 1
	j = 4
	defer fmt.Println(j)
	// 4 1
}
