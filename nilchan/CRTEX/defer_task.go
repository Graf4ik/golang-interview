package CRTEX

import "fmt"

// Что выведет программа?

func main() {
	a := 10
	defer func() {
		fmt.Println("call 0", a+10) // ← вычислится при выходе: a == 11
	}()
	defer fmt.Println("call 1", a+10) // ← a == 10, вычислится сразу
	a++
	fmt.Println("call 2", a)
}

// 11 20 21
