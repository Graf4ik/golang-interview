package main

import "fmt"

func main() {
	a()
}

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(x, y, z) // что выведет консоль
	// [0 1 2], [0 1 2 4], [0 1 2 4]
}
